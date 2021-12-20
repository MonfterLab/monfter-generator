package models

import (
	"github.com/astaxie/beego/orm"
	"monfter-generator/utils"
)

func MonfterTraitsIns() *MonfterTraits {
	return &MonfterTraits{}
}

func init() {
	orm.RegisterModel(new(MonfterTraits))
}

type MonfterTraits struct {
	Id          int    `orm:"column(id);pk"`
	TokenId     string `orm:"column(token_id);"`
	DisplayType string `orm:"column(display_type);"`
	TraitType   string `orm:"column(trait_type);"`
	Value       string `orm:"column(value);"`
	CreateTime  string `orm:"column(create_time)"`
}

func (m *MonfterTraits) TableName() string {
	return "monfter_traits"
}

func (m *MonfterTraits) Add() (int64, error) {
	o := orm.NewOrm()

	nowtime := utils.GetNowDateTime()

	m.CreateTime = nowtime

	id, err := o.Insert(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MonfterTraits) Save() (bool, error) {
	var (
		o     = orm.NewOrm()
		trait = MonfterTraits{Id: m.Id}
	)

	err := o.Read(&trait)
	if err != nil {
		return false, err
	}

	nowtime := utils.GetNowDateTime()
	m.CreateTime = nowtime

	_, err = o.Update(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MonfterTraits) Delete() (bool, error) {
	o := orm.NewOrm()

	_, err := o.Delete(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MonfterTraits) DeleteByTokenId(tokenId string) (int, error) {
	o := orm.NewOrm()

	num, err := o.QueryTable(m.TableName()).Filter("token_id", tokenId).Delete()
	if err != nil {
		return 0, err
	}

	return int(num), nil
}

func (m *MonfterTraits) Find(id int) (MonfterTraits, error) {
	var (
		o     = orm.NewOrm()
		trait = MonfterTraits{Id: id}
	)
	err := o.Read(&trait)
	if err != nil {
		return trait, err
	}
	return trait, nil
}

func (m *MonfterTraits) FindByTokenIdAndType(tokenId, traitType string) (MonfterTraits, error) {
	var (
		o     = orm.NewOrm()
		trait = MonfterTraits{TokenId: tokenId, TraitType: traitType}
	)
	err := o.Read(&trait, "token_id", "trait_type")
	if err != nil {
		return trait, err
	}
	return trait, nil
}

func (m *MonfterTraits) GetByTokenId(tokenId interface{}) ([]MonfterTraits, int64, error) {
	o := orm.NewOrm()
	list := make([]MonfterTraits, 0)

	num, err := o.QueryTable(m.TableName()).Filter("token_id", tokenId).All(&list)
	if err != nil {
		return list, 0, err
	}

	return list, num, nil
}
