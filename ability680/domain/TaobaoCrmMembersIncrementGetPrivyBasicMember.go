package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoCrmMembersIncrementGetPrivyBasicMember struct {
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

func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetOuid(v string) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.Ouid = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetItemNum(v int64) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.ItemNum = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetCloseTradeAmount(v string) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.CloseTradeAmount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetGroupIds(v string) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.GroupIds = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetStatus(v string) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.Status = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetRelationSource(v int64) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.RelationSource = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetTradeAmount(v string) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.TradeAmount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetGrade(v int64) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.Grade = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetCloseTradeCount(v int64) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.CloseTradeCount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetLastTradeTime(v util.LocalTime) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.LastTradeTime = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetTradeCount(v int64) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.TradeCount = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetBizOrderId(v int64) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.BizOrderId = &v
	return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyBasicMember) SetGradeName(v string) *TaobaoCrmMembersIncrementGetPrivyBasicMember {
	s.GradeName = &v
	return s
}
