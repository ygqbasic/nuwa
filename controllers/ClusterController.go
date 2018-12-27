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
	"github.com/ygqbasic/nuwa/runtime"
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

func (this *ClusterController) Detail() {

	Id, _ := this.GetInt(":clusterid", 0)
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
	this.Data["pageTitle"] = m.Name
	this.Data["showMoreQuery"] = true
	this.Data["clusterId"] = m.Id

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	this.setTpl()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headcssjs"] = "cluster/detail_headcssjs.html"
	this.LayoutSections["footerjs"] = "cluster/detail_footerjs.html"

	//页面里按钮权限控制
	this.Data["canEdit"] = this.checkActionAuthor("ClusterController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("ClusterController", "Delete")
}

// @Title RetrieveClusters
// @Description get all cluster
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.Cluster
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /retrieveclusters [get]
func (this *ClusterController) RetrieveClusters() {
	var params models.ClusterQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	data, total := models.RetrieveClusters(&params)

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

// @Title Save
// @Description add | update cluster info
// @Param   Name     path    string  true        "cluster name"
// @Param   Description     path    string  true        "cluster description"
// @Success 200 {object} models.Cluster
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /save [post]
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

func (this *ClusterController) install() {
	var err error
	op := &runtime.LaunchParameters{}

	if err = this.ParseForm(&op); err != nil {
		this.jsonResult(enums.JRCodeFailed, err.Error(), op)
		return
	}

	clusterID := this.GetString("cluster_id")
	id, err := strconv.Atoi(clusterID)
	o := orm.NewOrm()
	cluster := models.Cluster{Id: id}

	err = o.Read(&cluster)

	if err != nil {
		this.jsonResult(enums.JRCodeFailed, err.Error(), op)
		return
	}

	if err := runtime.StartOperate(&cluster, op); err != nil {
		this.jsonResult(enums.JRCodeFailed, err.Error(), op)
		return
	}

	cluster.State = models.Processing
	cluster.ChangeUser = this.curUser.RealName
	cluster.ChangeDate = time.Now()
	if _, err = o.Update(&cluster, "State", "ChangeUser", "ChangeDate"); err != nil {
		this.jsonResult(enums.JRCodeFailed, err.Error(), cluster)
		return
	}

	this.jsonResult(enums.JRCodeSucc, "stated", cluster.Id)
}
