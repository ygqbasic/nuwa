package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Host struct {
	Id          int       `form:"Id"`
	HostName    string    `orm:"size(50);unique" Json:"hostname"`
	Ip          string    `orm:"size(24)" Json:"ip"`
	Description string    `orm:"size(256)" Json:"description"`
	CreateUser  string    `orm:"column(createuser)" form:"CreateUser"`
	CreateDate  time.Time `orm:"column(createdate)" form:"CreateDate"`
	ChangeUser  string    `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate  time.Time `orm:"column(changedate);null" form:"ChangeDate"`
}

func init() {
	orm.RegisterModel(new(Host))
}

//角色与用户多对多关系表
func HostTBName() string {
	return "kube_host"
}

func (this *Host) TableName() string {
	return HostTBName()
}

type HostQueryParam struct {
	BaseQueryParam
	HostName string
	Ip       string //为空不查询，有值精确查询
}

func RetrieveHosts(clusterid int) ([]*ClusterHostRel, int64) {
	query := orm.NewOrm().QueryTable(ClusterHostRelTBName())
	data := make([]*ClusterHostRel, 0)

	query = query.Filter("Cluster__Id", clusterid)

	total, _ := query.Count()
	query.All(&data)

	return data, total
}

func HostPageList(params *HostQueryParam) ([]*Host, int64) {
	query := orm.NewOrm().QueryTable(HostTBName())
	data := make([]*Host, 0)

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

func HostDataList(params *HostQueryParam) []*Host {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := HostPageList(params)
	return data
}

//delete
func HostBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
