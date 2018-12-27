package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ClusterLog struct {
	Id        int       `form:"Id"`
	Now       time.Time `orm:"column(now)" form:"time"`
	Ip        string    `orm:"size(24)" Json:"ip"`
	TaskName  string    `orm:"size(24)" Json:"task_name"`
	TaskState string    `orm:"size(256)" Json:"task_state"`
	Stage     string    `orm:"size(256)" Json:"stage"`
	State     string    `orm:"size(256)" Json:"state"`
	Host      string    `orm:"size(256)" Json:"host"`
	Data      string    `orm:"size(256)" Json:"data"`

	Cluster    *Cluster  `orm:"rel(fk)"`
	CreateUser string    `orm:"column(createuser)" form:"CreateUser"`
	CreateDate time.Time `orm:"column(createdate)" form:"CreateDate"`
	ChangeUser string    `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate time.Time `orm:"column(changedate);null" form:"ChangeDate"`
}

func init() {
	orm.RegisterModel(new(ClusterLog))
}

//角色与用户多对多关系表
func ClusterLogTBName() string {
	return "kube_cluster_log"
}

func (this *ClusterLog) TableName() string {
	return ClusterLogTBName()
}

type ClusterLogQueryParam struct {
	BaseQueryParam
	HostName string
	Ip       string //为空不查询，有值精确查询
}

func RetrieveClusterLogs(clusterid int) ([]*ClusterLog, int64) {
	query := orm.NewOrm().QueryTable(ClusterLogTBName())
	data := make([]*ClusterLog, 0)

	query = query.Filter("Cluster__Id", clusterid)

	total, _ := query.Count()
	query.All(&data)

	return data, total
}

func ClusterLogPageList(params *ClusterLogQueryParam) ([]*ClusterLog, int64) {
	query := orm.NewOrm().QueryTable(ClusterLogTBName())
	data := make([]*ClusterLog, 0)

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

func ClusterLogDataList(params *ClusterLogQueryParam) []*ClusterLog {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := ClusterLogPageList(params)
	return data
}

//delete
func ClusterLogBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func ClusterLogsDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterLogTBName())
	num, err := query.Filter("ClusterId", id).Delete()
	return num, err
}
