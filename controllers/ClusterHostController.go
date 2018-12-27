package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/ygqbasic/nuwa/enums"
	"github.com/ygqbasic/nuwa/models"
)

type ClusterHostController struct {
	BaseController
}

func (this *ClusterHostController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor("DataGrid", "DataList")
}

func (this *ClusterHostController) Index() {
	this.Data["pageTitle"] = "主机管理"
	this.Data["showMoreQuery"] = true

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	this.setTpl()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headcssjs"] = "clusterhost/index_headcssjs.html"
	this.LayoutSections["footerjs"] = "clusterhost/index_footerjs.html"

	//页面里按钮权限控制
	this.Data["canEdit"] = this.checkActionAuthor("ClusterHostController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("ClusterHostController", "Delete")
}

func (this *ClusterHostController) RetrieveHosts() {
	Id, _ := this.GetInt(":cluster_id", 0)

	data, total := models.RetrieveClusterHosts(Id)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	//this.jsonResult(enums.JRCodeSucc, "", data)
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterHostController) DataCard() {
	var params models.ClusterHostQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.ClusterHostPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	//this.jsonResult(enums.JRCodeSucc, "", data)
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterHostController) DataGrid() {
	var params models.ClusterHostQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.ClusterHostPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterHostController) DataList() {
	var params = models.ClusterHostQueryParam{}
	data := models.ClusterHostDataList(&params)
	this.jsonResult(enums.JRCodeSucc, "", data)
}

func (this *ClusterHostController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	CId, _ := this.GetInt(":cluster_id", 0)
	Id, _ := this.GetInt(":id", 0)
	m := models.ClusterHost{Id: Id, Cluster: &models.Cluster{Id: CId}}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			this.pageError("数据无效，请刷新后重试" + string(CId))
		}
	} else {
		//m.State = "initial"
	}

	this.Data["m"] = m
	this.setTpl("clusterhost/edit.html", "shared/layout_pullbox.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footerjs"] = "clusterhost/edit_footerjs.html"
}

//add | update
func (this *ClusterHostController) Save() {
	var err error
	m := models.ClusterHost{}

	//获取form里的值
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	id := this.Input().Get("Id")
	m.Id, _ = strconv.Atoi(id)

	// m.HostName = this.GetString("HostName")
	// m.Ip = this.GetString("Ip")
	// m.Description = this.GetString("Description")
	clusterId, _ := this.GetInt("ClusterId")
	m.Cluster = &models.Cluster{Id: clusterId}

	o := orm.NewOrm()
	if m.Id == 0 {
		m.CreateUser = this.curUser.RealName
		m.CreateDate = time.Now()

		if _, err = o.Insert(&m); err == nil {
			this.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		m.ChangeUser = this.curUser.RealName
		m.ChangeDate = time.Now()

		if _, err = o.Update(&m, "HostName", "Description", "ChangeUser", "ChangeDate"); err == nil {
			this.jsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

func (this *ClusterHostController) Delete() {
	strs := this.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}

	if num, err := models.ClusterHostBatchDelete(ids); err == nil {
		this.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
