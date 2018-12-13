package models

import "github.com/astaxie/beego/orm"

type ClusterHostRel struct {
	Id      int
	Cluster *Cluster `orm:"rel(fk)"`  //外键
	Host    *Host    `orm:"rel(fk)" ` //外键
}

func init() {
	orm.RegisterModel(new(ClusterHostRel))
}

//角色与用户多对多关系表
func ClusterHostRelTBName() string {
	return "kube_cluster_host_rel"
}

func (this *ClusterHostRel) TableName() string {
	return ClusterHostRelTBName()
}
