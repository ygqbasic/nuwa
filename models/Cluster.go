package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	Initial    = "initial"
	Processing = "processing"
	Success    = "success"
	Failed     = "failed"
)

type Cluster struct {
	Id             int               `form:"Id"`
	Name           string            `orm:"size(256);unique" Json:"name"`
	Description    string            `orm:"size(256)" Json:"description"`
	State          string            `orm:"size(256)" Json:"state"`
	ClusterHostRel []*ClusterHostRel `orm:"reverse(many)" Json:"-"`
	CreateUser     string            `orm:"column(createuser)" form:"CreateUser"`
	CreateDate     time.Time         `orm:"column(createdate)" form:"CreateDate"`
	ChangeUser     string            `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate     time.Time         `orm:"column(changedate);null" form:"ChangeDate"`
}

func init() {
	orm.RegisterModel(new(Cluster))
}

//
func ClusterTBName() string {
	return "kube_cluster"
}

func (this *Cluster) TableName() string {
	return ClusterTBName()
}

type ClusterQueryParam struct {
	BaseQueryParam
	Name  string
	State string //为空不查询，有值精确查询
}

func RetrieveClusters(params *ClusterQueryParam) ([]*Cluster, int64) {
	query := orm.NewOrm().QueryTable(ClusterTBName())
	data := make([]*Cluster, 0)

	query = query.Filter("Name__istartswith", params.Name)
	query = query.Filter("State__istartswith", params.State)

	total, _ := query.Count()
	query.All(&data)

	return data, total
}

func ClusterPageList(params *ClusterQueryParam) ([]*Cluster, int64) {
	query := orm.NewOrm().QueryTable(ClusterTBName())
	data := make([]*Cluster, 0)

	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "State":
		sortorder = "State"
	}

	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("Name__istartswith", params.Name)
	query = query.Filter("State__istartswith", params.State)

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func ClusterDataList(params *ClusterQueryParam) []*Cluster {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := ClusterPageList(params)
	return data
}

//delete
func ClusterBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
