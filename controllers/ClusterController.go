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

type ClusterController struct {
	BaseController
}

func (this *ClusterController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor("DataGrid", "DataList")
}

func (this *ClusterController) Index() {
	this.Data["pageTitle"] = "集群管理"
	this.Data["showMoreQuery"] = true

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	this.setTpl()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headcssjs"] = "cluster/index_headcssjs.html"
	this.LayoutSections["footerjs"] = "cluster/index_footerjs.html"

	//页面里按钮权限控制
	this.Data["canEdit"] = this.checkActionAuthor("ClusterController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("ClusterController", "Delete")
}

func (this *ClusterController) DataCard() {
	var params models.ClusterQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.ClusterPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	//result["rows"] = data

	this.jsonResult(enums.JRCodeSucc, "", data)
	// this.Data["json"] = result
	// this.ServeJSON()
}
func (this *ClusterController) DataGrid() {
	var params models.ClusterQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.ClusterPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ClusterController) DataList() {
	var params = models.ClusterQueryParam{}
	data := models.ClusterDataList(&params)
	this.jsonResult(enums.JRCodeSucc, "", data)
}

func (this *ClusterController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt(":id", 0)
	m := models.Cluster{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			this.pageError("数据无效，请刷新后重试")
		}
	} else {
		m.State = "initial"
	}

	this.Data["m"] = m
	this.setTpl("cluster/edit.html", "shared/layout_pullbox.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footerjs"] = "cluster/edit_footerjs.html"
}

//add | update
func (this *ClusterController) Save() {
	var err error
	m := models.Cluster{}

	//获取form里的值
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	id := this.Input().Get("Id")
	m.Id, _ = strconv.Atoi(id)

	m.Name = this.GetString("Name")
	m.Description = this.GetString("Description")

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

		if _, err = o.Update(&m, "Name", "Description", "ChangeUser", "ChangeDate"); err == nil {
			this.jsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

func (this *ClusterController) Delete() {
	strs := this.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}

	if num, err := models.ClusterBatchDelete(ids); err == nil {
		this.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
