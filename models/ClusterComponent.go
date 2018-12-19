package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ClusterComponent struct {
	Id            int       `form:"Id"`
	ComponentName string    `orm:"size(24)" Json:"ComponentName"`
	Component     string    `orm:"size(246)" Json:"Component"`
	Description   string    `orm:"size(256)" Json:"description"`
	Cluster       *Cluster  `orm:"rel(fk)"`
	CreateUser    string    `orm:"column(createuser)" form:"CreateUser"`
	CreateDate    time.Time `orm:"column(createdate)" form:"CreateDate"`
	ChangeUser    string    `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate    time.Time `orm:"column(changedate);null" form:"ChangeDate"`
}

func init() {
	orm.RegisterModel(new(ClusterComponent))
}

//角色与用户多对多关系表
func ClusterComponentTBName() string {
	return "kube_cluster_component"
}

func (this *ClusterComponent) TableName() string {
	return ClusterComponentTBName()
}

type ClusterComponentQueryParam struct {
	BaseQueryParam
	HostName string
	Ip       string //为空不查询，有值精确查询
}

func RetrieveClusterComponents(clusterid int) ([]*ClusterComponent, int64) {
	query := orm.NewOrm().QueryTable(ClusterComponentTBName())
	data := make([]*ClusterComponent, 0)

	query = query.Filter("Cluster__Id", clusterid)

	total, _ := query.Count()
	query.All(&data)

	return data, total
}

func ClusterComponentPageList(params *ClusterComponentQueryParam) ([]*ClusterComponent, int64) {
	query := orm.NewOrm().QueryTable(ClusterComponentTBName())
	data := make([]*ClusterComponent, 0)

	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Ip":
		sortorder = "Ip"
	}

	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("HostName__istartswith", params.HostName)
	query = query.Filter("Ip__istartswith", params.Ip)

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func ClusterComponentDataList(params *ClusterComponentQueryParam) []*ClusterComponent {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := ClusterComponentPageList(params)
	return data
}

//delete
func ClusterComponentBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterComponentTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

type MetaComponent struct {
	Id       int                    `json:"id"`
	Name     string                 `json:"name"`
	Version  string                 `json:"version"`
	Property map[string]interface{} `json:"properties"`
}

type Component struct {
	MetaComponent
	Hosts map[string][]string `json:"hosts"`
}
