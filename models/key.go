package models

import (
	"github.com/astaxie/beego/orm"
)

func MonfterKeyIns() *MonfterKey {
	return &MonfterKey{}
}

func init() {
	orm.RegisterModel(new(MonfterKey))
}

type MonfterKey struct {
	Id      int    `orm:"column(id);pk"`
	Key     string `orm:"column(key);"`
	TokenId string `orm:"column(token_id);"`
}

func (m *MonfterKey) TableName() string {
	return "monfter_key"
}

func (m *MonfterKey) Add() (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MonfterKey) Save(MonfterKeyId, ownerId int) (bool, error) {
	var (
		o          = orm.NewOrm()
		MonfterKey = MonfterKey{Id: MonfterKeyId}
	)

	err := o.Read(&MonfterKey)
	if err != nil {
		return false, err
	}

	_, err = o.Update(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MonfterKey) DeleteByTokenId(tokenId string) (int, error) {
	o := orm.NewOrm()

	num, err := o.QueryTable(m.TableName()).Filter("token_id", tokenId).Delete()
	if err != nil {
		return 0, err
	}

	return int(num), nil
}

func (m *MonfterKey) FindByKey(key string) (MonfterKey, error) {
	var (
		o   = orm.NewOrm()
		row = MonfterKey{Key: key}
	)
	err := o.Read(&row, "key")
	if err != nil {
		return row, err
	}
	return row, nil
}
