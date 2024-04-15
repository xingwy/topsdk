package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTradeOrderskuUpdateOrder struct {
	/*
	   子订单编号     */
	Oid *int64 `json:"oid,omitempty" `

	/*
	   订单修改时间，目前只有taobao.trade.ordersku.update会返回此字段。     */
	Modified *util.LocalTime `json:"modified,omitempty" `
}

func (s *TaobaoTradeOrderskuUpdateOrder) SetOid(v int64) *TaobaoTradeOrderskuUpdateOrder {
	s.Oid = &v
	return s
}
func (s *TaobaoTradeOrderskuUpdateOrder) SetModified(v util.LocalTime) *TaobaoTradeOrderskuUpdateOrder {
	s.Modified = &v
	return s
}
