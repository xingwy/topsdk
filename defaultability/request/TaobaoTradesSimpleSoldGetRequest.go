package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTradesSimpleSoldGetRequest struct {
	/*
	   买家的openId     */
	BuyerOpenId *string `json:"buyer_open_id,omitempty" required:"false" `
	/*
	   需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0,处方药未审核     */
	Fields *[]string `json:"fields" required:"true" `
	/*
	   查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss     */
	StartCreated *util.LocalTime `json:"start_created,omitempty" required:"false" `
	/*
	   查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss     */
	EndCreated *util.LocalTime `json:"end_created,omitempty" required:"false" `
	/*
	   交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。     */
	Status *string `json:"status,omitempty" required:"false" `
	/*
	   买家昵称     */
	BuyerNick *string `json:"buyer_nick,omitempty" required:"false" `
	/*
	   交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade,auto_delivery,ec,cod,step 这5 种的交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>可选值：fixed(一口价)auction(拍卖)guarantee_trade(一口价、拍卖)step(分阶段付款，万人团，阶梯团订单）independent_simple_trade(旺店入门版交易)independent_shop_trade(旺店标准版交易)auto_delivery(自动发货)ec(直冲)cod(货到付款)game_equipment(游戏装备)shopex_trade(ShopEX交易)netcn_trade(万网交易)external_trade(统一外部交易)instant_trade (即时到账)b2c_cod(大商家货到付款)hotel_trade(酒店类型交易)super_market_trade(商超交易)super_market_cod_trade(商超货到付款交易)taohua(淘花网交易类型）waimai(外卖交易类型）o2o_offlinetrade（O2O交易）nopaid（即时到帐/趣味猜交易类型）step (万人团) eticket(电子凭证) tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。pre_auth_type(预授权0元购) lazada（获取lazada订单类型）     */
	Type *string `json:"type,omitempty" required:"false" `
	/*
	   可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型     */
	ExtType *string `json:"ext_type,omitempty" required:"false" `
	/*
	   评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评)     */
	RateStatus *string `json:"rate_status,omitempty" required:"false" `
	/*
	   卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）     */
	Tag *string `json:"tag,omitempty" required:"false" `
	/*
	   页码。默认值:1，可填整数，通过传入page_no来控制获取的页数，总页数=total_results÷page_size defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   每页条数。 默认值:40;最大值:100，可填整数。通过page_no和page_size组合多次调用实现翻页获取全量数据。 defalutValue��40    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。 defalutValue��false    */
	UseHasNext *bool `json:"use_has_next,omitempty" required:"false" `
}

func (s *TaobaoTradesSimpleSoldGetRequest) SetBuyerOpenId(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.BuyerOpenId = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetFields(v []string) *TaobaoTradesSimpleSoldGetRequest {
	s.Fields = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetStartCreated(v util.LocalTime) *TaobaoTradesSimpleSoldGetRequest {
	s.StartCreated = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetEndCreated(v util.LocalTime) *TaobaoTradesSimpleSoldGetRequest {
	s.EndCreated = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetStatus(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.Status = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetBuyerNick(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.BuyerNick = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetType(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.Type = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetExtType(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.ExtType = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetRateStatus(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.RateStatus = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetTag(v string) *TaobaoTradesSimpleSoldGetRequest {
	s.Tag = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetPageNo(v int64) *TaobaoTradesSimpleSoldGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetPageSize(v int64) *TaobaoTradesSimpleSoldGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTradesSimpleSoldGetRequest) SetUseHasNext(v bool) *TaobaoTradesSimpleSoldGetRequest {
	s.UseHasNext = &v
	return s
}

func (req *TaobaoTradesSimpleSoldGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BuyerOpenId != nil {
		paramMap["buyer_open_id"] = *req.BuyerOpenId
	}
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	if req.StartCreated != nil {
		paramMap["start_created"] = *req.StartCreated
	}
	if req.EndCreated != nil {
		paramMap["end_created"] = *req.EndCreated
	}
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	if req.BuyerNick != nil {
		paramMap["buyer_nick"] = *req.BuyerNick
	}
	if req.Type != nil {
		paramMap["type"] = *req.Type
	}
	if req.ExtType != nil {
		paramMap["ext_type"] = *req.ExtType
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
	return paramMap
}

func (req *TaobaoTradesSimpleSoldGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
