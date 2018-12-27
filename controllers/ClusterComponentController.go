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

type ClusterComponentController struct {
	BaseController
}

func (this *ClusterComponentController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor("DataGrid", "DataList")
}

func (this *ClusterComponentController) Index() {
	this.Data["pageTitle"] = "主机管理"
	this.Data["showMoreQuery"] = true

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	this.setTpl()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headcssjs"] = "clustercomponent/index_headcssjs.html"
	this.LayoutSections["footerjs"] = "clustercomponent/index_footerjs.html"

	//页面里按钮权限控制
	this.Data["canEdit"] = this.checkActionAuthor("ClusterComponentController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("ClusterComponentController", "Delete")
}

func (this *ClusterComponentController) RetrieveComponents() {
	Id, _ := this.GetInt(":cluster_id", 0)

	data, total := models.RetrieveClusterComponents(Id)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	//this.jsonResult(enums.JRCodeSucc, "", data)
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterComponentController) DataCard() {
	var params models.ClusterComponentQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.ClusterComponentPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	//this.jsonResult(enums.JRCodeSucc, "", data)
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterComponentController) DataGrid() {
	var params models.ClusterComponentQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.ClusterComponentPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterComponentController) DataList() {
	var params = models.ClusterComponentQueryParam{}
	data := models.ClusterComponentDataList(&params)
	this.jsonResult(enums.JRCodeSucc, "", data)
}

func (this *ClusterComponentController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}
	CId, _ := this.GetInt(":cluster_id", 0)
	Id, _ := this.GetInt(":id", 0)
	m := models.ClusterComponent{Id: Id, Cluster: &models.Cluster{Id: CId}}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			this.pageError("数据无效，请刷新后重试")
		}
	} else {
		//m.State = "initial"
	}

	this.Data["m"] = m
	this.setTpl("clustercomponent/edit.html", "shared/layout_pullbox.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footerjs"] = "clustercomponent/edit_footerjs.html"
}

//add | update
func (this *ClusterComponentController) Save() {
	var err error
	m := models.ClusterComponent{}

	//获取form里的值
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	id := this.Input().Get("Id")
	m.Id, _ = strconv.Atoi(id)

	clusterId, _ := this.GetInt("ClusterId")
	m.Cluster = &models.Cluster{Id: clusterId}

	m.Hosts = strings.Join(this.GetStrings("Hosts"), ",")

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

		if _, err = o.Update(&m, "Version", "ComponentName", "Description", "Properties", "Hosts", "ChangeUser", "ChangeDate"); err == nil {
			this.jsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

func (this *ClusterComponentController) Delete() {
	strs := this.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}

	if num, err := models.ClusterComponentBatchDelete(ids); err == nil {
		this.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
