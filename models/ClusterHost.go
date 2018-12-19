package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ClusterHost struct {
	Id          int       `form:"Id"`
	HostName    string    `orm:"size(50);unique" Json:"hostname"`
	Ip          string    `orm:"size(24)" Json:"ip"`
	Description string    `orm:"size(256)" Json:"description"`
	Cluster     *Cluster  `orm:"rel(fk)"`
	CreateUser  string    `orm:"column(createuser)" form:"CreateUser"`
	CreateDate  time.Time `orm:"column(createdate)" form:"CreateDate"`
	ChangeUser  string    `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate  time.Time `orm:"column(changedate);null" form:"ChangeDate"`
}

func init() {
	orm.RegisterModel(new(ClusterHost))
}

//角色与用户多对多关系表
func ClusterHostTBName() string {
	return "kube_cluster_host"
}

func (this *ClusterHost) TableName() string {
	return ClusterHostTBName()
}

type ClusterHostQueryParam struct {
	BaseQueryParam
	HostName string
	Ip       string //为空不查询，有值精确查询
}

func RetrieveClusterHosts(clusterid int) ([]*ClusterHost, int64) {
	query := orm.NewOrm().QueryTable(ClusterHostTBName())
	data := make([]*ClusterHost, 0)

	query = query.Filter("Cluster__Id", clusterid)

	total, _ := query.Count()
	query.All(&data)

	return data, total
}

func RetrieveClusterHost(
	clusterID int,
	hostID string,
) *ClusterHost {

	query := orm.NewOrm().QueryTable(ClusterHostTBName())
	data := make([]*ClusterHost, 0)

	query = query.Filter("Cluster__Id", clusterID)
	query = query.Filter("Id", hostID)

	total, _ := query.Count()
	query.All(&data)

	if total < 1 {
		return nil // ,fmt.Errorf("can't find cluster %s host %s", clusterID, hostID)
	}

	return data[0]
}

func ClusterHostPageList(params *ClusterHostQueryParam) ([]*ClusterHost, int64) {
	query := orm.NewOrm().QueryTable(ClusterHostTBName())
	data := make([]*ClusterHost, 0)

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

func ClusterHostDataList(params *ClusterHostQueryParam) []*ClusterHost {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := ClusterHostPageList(params)
	return data
}

//delete
func ClusterHostBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

type Host struct {
	Id          string `json:"id"`
	HostName    string `json:"hostname"`
	Ip          string `json:"ip"`
	Description string `json:"description"`
}
