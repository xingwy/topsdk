package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoRefundGetRefundRemindTimeout struct {
	/*
	   是否存在超时。可选值:true(是),false(否)     */
	ExistTimeout *bool `json:"exist_timeout,omitempty" `

	/*
	   提醒的类型（退款详情中提示信息的类型）     */
	RemindType *int64 `json:"remind_type,omitempty" `

	/*
	   超时时间。格式:yyyy-MM-dd HH:mm:ss     */
	Timeout *util.LocalTime `json:"timeout,omitempty" `
}

func (s *TaobaoRefundGetRefundRemindTimeout) SetExistTimeout(v bool) *TaobaoRefundGetRefundRemindTimeout {
	s.ExistTimeout = &v
	return s
}
func (s *TaobaoRefundGetRefundRemindTimeout) SetRemindType(v int64) *TaobaoRefundGetRefundRemindTimeout {
	s.RemindType = &v
	return s
}
func (s *TaobaoRefundGetRefundRemindTimeout) SetTimeout(v util.LocalTime) *TaobaoRefundGetRefundRemindTimeout {
	s.Timeout = &v
	return s
}
