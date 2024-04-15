package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTradesSimpleSoldIncrementGetRequest struct {
	/*
	   需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0为处方药未审核状态     */
	Fields *[]string `json:"fields" required:"true" `
	/*
	   查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss     */
	StartModified *util.LocalTime `json:"start_modified" required:"true" `
	/*
	   查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。     */
	EndModified *util.LocalTime `json:"end_modified" required:"true" `
	/*
	   交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。     */
	Status *string `json:"status,omitempty" required:"false" `
	/*
	   交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。     */
	Type *string `json:"type,omitempty" required:"false" `
	/*
	   可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型     */
	ExtType *string `json:"ext_type,omitempty" required:"false" `
	/*
	   买家昵称     */
	BuyerNick *string `json:"buyer_nick,omitempty" required:"false" `
	/*
	   评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评)     */
	RateStatus *string `json:"rate_status,omitempty" required:"false" `
	/*
	   卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）     */
	Tag *string `json:"tag,omitempty" required:"false" `
	/*
	   页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 defalutValue��40    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 defalutValue��false    */
	UseHasNext *bool `json:"use_has_next,omitempty" required:"false" `
	/*
	   用户openId     */
	BuyerOpenUid *string `json:"buyer_open_uid,omitempty" required:"false" `
}

func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetFields(v []string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.Fields = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetStartModified(v util.LocalTime) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.StartModified = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetEndModified(v util.LocalTime) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.EndModified = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetStatus(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.Status = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetType(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.Type = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetExtType(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.ExtType = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetBuyerNick(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.BuyerNick = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetRateStatus(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.RateStatus = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetTag(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.Tag = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetPageNo(v int64) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetPageSize(v int64) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetUseHasNext(v bool) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.UseHasNext = &v
	return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetRequest) SetBuyerOpenUid(v string) *TaobaoTradesSimpleSoldIncrementGetRequest {
	s.BuyerOpenUid = &v
	return s
}

func (req *TaobaoTradesSimpleSoldIncrementGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	if req.StartModified != nil {
		paramMap["start_modified"] = *req.StartModified
	}
	if req.EndModified != nil {
		paramMap["end_modified"] = *req.EndModified
	}
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	if req.Type != nil {
		paramMap["type"] = *req.Type
	}
	if req.ExtType != nil {
		paramMap["ext_type"] = *req.ExtType
	}
	if req.BuyerNick != nil {
		paramMap["buyer_nick"] = *req.BuyerNick
	}
	if req.RateStatus != nil {
		paramMap["rate_status"] = *req.RateStatus
	}
	if req.Tag != nil {
		paramMap["tag"] = *req.Tag
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.UseHasNext != nil {
		paramMap["use_has_next"] = *req.UseHasNext
	}
	if req.BuyerOpenUid != nil {
		paramMap["buyer_open_uid"] = *req.BuyerOpenUid
	}
	return paramMap
}

func (req *TaobaoTradesSimpleSoldIncrementGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
