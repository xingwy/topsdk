package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoCrmMembersIncrementGetBasicMember struct {
	/*
	   会员昵称     */
	BuyerNick *string `json:"buyer_nick,omitempty" `

	/*
	   open_uid     */
	Ouid *string `json:"ouid,omitempty" `

	/*
	   购买的宝贝件数     */
	ItemNum *int64 `json:"item_num,omitempty" `

	/*
	   交易关闭金额     */
	CloseTradeAmount *string `json:"close_trade_amount,omitempty" `

	/*
	   分组的id集合字符串     */
	GroupIds *string `json:"group_ids,omitempty" `

	/*
	   显示会员的状态，normal正常，blacklist黑名单     */
	Status *string `json:"status,omitempty" `

	/*
	   关系来源，1交易成功，2未交易成功 ,3 卖家主动吸纳     */
	RelationSource *int64 `json:"relation_source,omitempty" `

	/*
	   交易的金额     */
	TradeAmount *string `json:"trade_amount,omitempty" `

	/*
	   会员等级，0未非会员，其余对应等级名称取grade_name     */
	Grade *int64 `json:"grade,omitempty" `

	/*
	   交易关闭的笔数     */
	CloseTradeCount *int64 `json:"close_trade_count,omitempty" `

	/*
	   最后交易的日期     */
	LastTradeTime *util.LocalTime `json:"last_trade_time,omitempty" `

	/*
	   交易成功的次数     */
	TradeCount *int64 `json:"trade_count,omitempty" `

	/*
	   最后一次交易的订单号.注:该字段从2014.4.23之后不再返回.     */
	BizOrderId *int64 `json:"biz_order_id,omitempty" `

	/*
	   等级名称     */
	GradeName *string `json:"grade_name,omitempty" `
}

func (s *TaobaoCrmMembersIncrementGetBasicMember) SetBuyerNick(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.BuyerNick = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetOuid(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.Ouid = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetItemNum(v int64) *TaobaoCrmMembersIncrementGetBasicMember {
	s.ItemNum = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetCloseTradeAmount(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.CloseTradeAmount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetGroupIds(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.GroupIds = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetStatus(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.Status = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetRelationSource(v int64) *TaobaoCrmMembersIncrementGetBasicMember {
	s.RelationSource = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetTradeAmount(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.TradeAmount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetGrade(v int64) *TaobaoCrmMembersIncrementGetBasicMember {
	s.Grade = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetCloseTradeCount(v int64) *TaobaoCrmMembersIncrementGetBasicMember {
	s.CloseTradeCount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetLastTradeTime(v util.LocalTime) *TaobaoCrmMembersIncrementGetBasicMember {
	s.LastTradeTime = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetTradeCount(v int64) *TaobaoCrmMembersIncrementGetBasicMember {
	s.TradeCount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetBizOrderId(v int64) *TaobaoCrmMembersIncrementGetBasicMember {
	s.BizOrderId = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetBasicMember) SetGradeName(v string) *TaobaoCrmMembersIncrementGetBasicMember {
	s.GradeName = &v
	return s
}
