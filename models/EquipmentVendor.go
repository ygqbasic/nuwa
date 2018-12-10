package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type EquipmentVendor struct {
	Id         int       `orm:"column(id)" form:"Id"`
	VendorNO   string    `orm:"column(vendor_no)" form:"VendorNO"`
	VendorDesc string    `orm:"column(vendor_desc)" form:"VendorDesc"`
	Used       int       `orm:"column(tag)" form:"Used"`
	CreateUser string    `orm:"column(createuser)" form:"CreateUser"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime);column(createdate)" form:"CreateDate"`
	ChangeUser string    `orm:"column(changeuser)" form:"ChangeUser"`
	ChangeDate time.Time `orm:"auto_now;type(datetime);column(changedate)" form:"ChangeDate"`
}

type EquipmentVendorQueryParam struct {
	BaseQueryParam
	VendorNO   string
	VendorDesc string
	Used       string
}

func init() {
	orm.RegisterModel(new(EquipmentVendor))
}

//设备供应商
func EquipmentVendorTBName() string {
	return "equipment_vendor"
}

func EquipmentVendorPageList(params *EquipmentVendorQueryParam) ([]*EquipmentVendor, int64) {
	query := orm.NewOrm().QueryTable(EquipmentVendorTBName())
	data := make([]*EquipmentVendor, 0)

	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "VendorNO":
		sortorder = "vendor_no"
	case "VendorDesc":
		sortorder = "vendor_desc"
	case "Used":
		sortorder = "tag"
	}

	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("VendorNO__istartswith", params.VendorNO)
	query = query.Filter("VendorDesc__istartswith", params.VendorDesc)
	query = query.Filter("tag__istartswith", params.Used)

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func EquipmentVendorDataList(params *EquipmentVendorQueryParam) [] *EquipmentVendor {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := EquipmentVendorPageList(params)
	return data
}

func EquipmentVendorBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(EquipmentVendorTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func EquipmentVendorOne(id int) (*EquipmentVendor, error) {
	o := orm.NewOrm()
	m := EquipmentVendor{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (this *EquipmentVendor) TableName() string {
	return EquipmentVendorTBName()
}
