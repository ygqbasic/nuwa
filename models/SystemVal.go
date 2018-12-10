package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type SystemVal struct {
	Id         int       `orm:"column(id)" form:"Id"`
	Code       string    `orm:"column(code)" form:"Code"`
	Desc       string    `orm:"column(desc)" form:"Desc"`
	Value      string    `orm:"column(value)" form:"Value"`
	UpLimit    string    `orm:"column(uplimit)" form:"UpLimit"`
	Step       string    `orm:"column(step)" form:"Step"`
	Used       int       `orm:"column(tag)" form:"Used"`
	CreateUser string    `orm:"column(createuser)" form:"CreateUser"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime);column(createdate)" form:"CreateDate"`
	ChangeUser string    `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate time.Time `orm:"auto_now;type(datetime);column(changedate)" form:"ChangeDate"`
}

type SystemValQueryParam struct {
	BaseQueryParam
	Code string
	Desc string
	Used string //为空不查询，有值精确查询
}

func init() {
	orm.RegisterModel(new(SystemVal))
}

func SystemValTBName() string {
	return "sys_val"
}

func (this *SystemVal) TableName() string {
	return SystemValTBName()
}

func SystemValPageList(params *SystemValQueryParam) ([]*SystemVal, int64) {
	query := orm.NewOrm().QueryTable(SystemValTBName())
	data := make([]*SystemVal, 0)

	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Used":
		sortorder = "tag"
	}

	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("code__contains", params.Code)
	query = query.Filter("desc__contains", params.Desc)
	query = query.Filter("tag__istartswith", params.Used)

	total, _ := query.Count()

	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func SystemValDataList(params *SystemValQueryParam) []*SystemVal {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := SystemValPageList(params)
	return data
}

func SystemValBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(SystemValTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func SystemValOne(id int) (*SystemVal, error) {
	o := orm.NewOrm()
	m := SystemVal{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
