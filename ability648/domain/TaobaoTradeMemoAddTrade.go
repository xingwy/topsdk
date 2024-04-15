package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTradeMemoAddTrade struct {
	/*
	   交易编号 (父订单的交易编号)     */
	Tid *int64 `json:"tid,omitempty" `

	/*
	   交易创建时间。格式:yyyy-MM-dd HH:mm:ss     */
	Created *util.LocalTime `json:"created,omitempty" `
}

func (s *TaobaoTradeMemoAddTrade) SetTid(v int64) *TaobaoTradeMemoAddTrade {
	s.Tid = &v
	return s
}
func (s *TaobaoTradeMemoAddTrade) SetCreated(v util.LocalTime) *TaobaoTradeMemoAddTrade {
	s.Created = &v
	return s
}
