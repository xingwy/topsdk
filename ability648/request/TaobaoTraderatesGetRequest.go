package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTraderatesGetRequest struct {
	/*
	   需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔     */
	Fields *[]string `json:"fields" required:"true" `
	/*
	   评价类型。可选值:get(得到),give(给出)     */
	RateType *string `json:"rate_type" required:"true" `
	/*
	   评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。     */
	Role *string `json:"role" required:"true" `
	/*
	   评价结果。可选值:good(好评),neutral(中评),bad(差评)     */
	Result *string `json:"result,omitempty" required:"false" `
	/*
	   页码。取值范围:大于零的整数最大限制为200; 默认值:1     */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   每页获取条数。默认值40，最小值1，最大值150。 defalutValue��40    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   评价对方昵称     */
	PeerNick *string `json:"peer_nick,omitempty" required:"false" `
	/*
	   评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。     */
	StartDate *util.LocalTime `json:"start_date,omitempty" required:"false" `
	/*
	   评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。     */
	EndDate *util.LocalTime `json:"end_date,omitempty" required:"false" `
	/*
	   交易订单id，可以是父订单id号，也可以是子订单id号     */
	Tid *int64 `json:"tid,omitempty" required:"false" `
	/*
	   是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。 defalutValue��false    */
	UseHasNext *bool `json:"use_has_next,omitempty" required:"false" `
	/*
	   商品的数字ID     */
	NumIid *int64 `json:"num_iid,omitempty" required:"false" `
	/*
	   ouid     */
	Ouid *string `json:"ouid,omitempty" required:"false" `
}

func (s *TaobaoTraderatesGetRequest) SetFields(v []string) *TaobaoTraderatesGetRequest {
	s.Fields = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetRateType(v string) *TaobaoTraderatesGetRequest {
	s.RateType = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetRole(v string) *TaobaoTraderatesGetRequest {
	s.Role = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetResult(v string) *TaobaoTraderatesGetRequest {
	s.Result = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetPageNo(v int64) *TaobaoTraderatesGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetPageSize(v int64) *TaobaoTraderatesGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetPeerNick(v string) *TaobaoTraderatesGetRequest {
	s.PeerNick = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetStartDate(v util.LocalTime) *TaobaoTraderatesGetRequest {
	s.StartDate = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetEndDate(v util.LocalTime) *TaobaoTraderatesGetRequest {
	s.EndDate = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetTid(v int64) *TaobaoTraderatesGetRequest {
	s.Tid = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetUseHasNext(v bool) *TaobaoTraderatesGetRequest {
	s.UseHasNext = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetNumIid(v int64) *TaobaoTraderatesGetRequest {
	s.NumIid = &v
	return s
}
func (s *TaobaoTraderatesGetRequest) SetOuid(v string) *TaobaoTraderatesGetRequest {
	s.Ouid = &v
	return s
}

func (req *TaobaoTraderatesGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	if req.RateType != nil {
		paramMap["rate_type"] = *req.RateType
	}
	if req.Role != nil {
		paramMap["role"] = *req.Role
	}
	if req.Result != nil {
		paramMap["result"] = *req.Result
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PeerNick != nil {
		paramMap["peer_nick"] = *req.PeerNick
	}
	if req.StartDate != nil {
		paramMap["start_date"] = *req.StartDate
	}
	if req.EndDate != nil {
		paramMap["end_date"] = *req.EndDate
	}
	if req.Tid != nil {
		paramMap["tid"] = *req.Tid
	}
	if req.UseHasNext != nil {
		paramMap["use_has_next"] = *req.UseHasNext
	}
	if req.NumIid != nil {
		paramMap["num_iid"] = *req.NumIid
	}
	if req.Ouid != nil {
		paramMap["ouid"] = *req.Ouid
	}
	return paramMap
}

func (req *TaobaoTraderatesGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
