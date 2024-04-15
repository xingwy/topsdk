package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoFuwuSpBillreordAddBillRecordDto struct {
	/*
	   appkey     */
	Appkey *string `json:"appkey,omitempty" `

	/*
	   备用字段     */
	Extend1 *string `json:"extend1,omitempty" `

	/*
	   备用字段     */
	Extend10 *string `json:"extend10,omitempty" `

	/*
	   备用字段     */
	Extend2 *string `json:"extend2,omitempty" `

	/*
	   备用字段     */
	Extend3 *string `json:"extend3,omitempty" `

	/*
	   备用字段     */
	Extend4 *string `json:"extend4,omitempty" `

	/*
	   备用字段     */
	Extend5 *string `json:"extend5,omitempty" `

	/*
	   备用字段     */
	Extend6 *string `json:"extend6,omitempty" `

	/*
	   备用字段     */
	Extend7 *string `json:"extend7,omitempty" `

	/*
	   备用字段     */
	Extend8 *string `json:"extend8,omitempty" `

	/*
	   备用字段     */
	Extend9 *string `json:"extend9,omitempty" `

	/*
	   金额     */
	Fee *int64 `json:"fee,omitempty" `

	/*
	   卖家nick     */
	Nick *string `json:"nick,omitempty" `

	/*
	   服务市场有效订单ID     */
	OrderId *int64 `json:"order_id,omitempty" `

	/*
	   外部确认账单ID     */
	OutConfirmId *string `json:"out_confirm_id,omitempty" `

	/*
	   外部订单ID     */
	OutOrderId *string `json:"out_order_id,omitempty" `

	/*
	   记录产生时间     */
	StartDate *util.LocalTime `json:"start_date,omitempty" `

	/*
	   状态：1成功、2失败     */
	Status *int64 `json:"status,omitempty" `

	/*
	   目标号码     */
	TargetNo *string `json:"target_no,omitempty" `

	/*
	   账单分类：1短信     */
	Type *int64 `json:"type,omitempty" `
}

func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetAppkey(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Appkey = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend1(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend1 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend10(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend10 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend2(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend2 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend3(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend3 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend4(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend4 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend5(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend5 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend6(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend6 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend7(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend7 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend8(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend8 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetExtend9(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Extend9 = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetFee(v int64) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Fee = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetNick(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Nick = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetOrderId(v int64) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.OrderId = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetOutConfirmId(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.OutConfirmId = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetOutOrderId(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.OutOrderId = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetStartDate(v util.LocalTime) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.StartDate = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetStatus(v int64) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Status = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetTargetNo(v string) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.TargetNo = &v
	return s
}
func (s *TaobaoFuwuSpBillreordAddBillRecordDto) SetType(v int64) *TaobaoFuwuSpBillreordAddBillRecordDto {
	s.Type = &v
	return s
}
