package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto struct {
	/*
	   应到达日期     */
	ScDate *util.LocalTime `json:"sc_date,omitempty" `

	/*
	   子交易单号     */
	SubTradeIds *[]string `json:"sub_trade_ids,omitempty" `

	/*
	   交易号     */
	TradeId *string `json:"trade_id,omitempty" `

	/*
	   卖家Id     */
	SellerId *int64 `json:"seller_id,omitempty" `

	/*
	   收货人电话     */
	ReceiverMobile *string `json:"receiver_mobile,omitempty" `

	/*
	   改约日期     */
	OsDate *util.LocalTime `json:"os_date,omitempty" `

	/*
	   收货人姓名     */
	ReceiverName *string `json:"receiver_name,omitempty" `

	/*
	   扩展字段     */
	Feature *string `json:"feature,omitempty" `

	/*
	   外部订单号     */
	OutOrderCode *string `json:"out_order_code,omitempty" `

	/*
	   收货人地址     */
	ReceiverAddress *string `json:"receiver_address,omitempty" `
}

func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetScDate(v util.LocalTime) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.ScDate = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetSubTradeIds(v []string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.SubTradeIds = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetTradeId(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.TradeId = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetSellerId(v int64) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.SellerId = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetReceiverMobile(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.ReceiverMobile = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetOsDate(v util.LocalTime) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.OsDate = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetReceiverName(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.ReceiverName = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetFeature(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.Feature = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetOutOrderCode(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.OutOrderCode = &v
	return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) SetReceiverAddress(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto {
	s.ReceiverAddress = &v
	return s
}
