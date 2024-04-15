package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTradePostageUpdateTrade struct {
	/*
	   交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss     */
	Modified *util.LocalTime `json:"modified,omitempty" `

	/*
	   邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	PostFee *string `json:"post_fee,omitempty" `

	/*
	   商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	TotalFee *string `json:"total_fee,omitempty" `

	/*
	   实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	Payment *string `json:"payment,omitempty" `
}

func (s *TaobaoTradePostageUpdateTrade) SetModified(v util.LocalTime) *TaobaoTradePostageUpdateTrade {
	s.Modified = &v
	return s
}
func (s *TaobaoTradePostageUpdateTrade) SetPostFee(v string) *TaobaoTradePostageUpdateTrade {
	s.PostFee = &v
	return s
}
func (s *TaobaoTradePostageUpdateTrade) SetTotalFee(v string) *TaobaoTradePostageUpdateTrade {
	s.TotalFee = &v
	return s
}
func (s *TaobaoTradePostageUpdateTrade) SetPayment(v string) *TaobaoTradePostageUpdateTrade {
	s.Payment = &v
	return s
}
