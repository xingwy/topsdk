package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoRefundsApplyGetRequest struct {
	/*
	   需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee     */
	Fields *[]string `json:"fields" required:"true" `
	/*
	   卖家昵称     */
	SellerNick *string `json:"seller_nick,omitempty" required:"false" `
	/*
		        退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。
		WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意)
		WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货)
		WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货)
		SELLER_REFUSE_BUYER(卖家拒绝退款)
		CLOSED(退款关闭)
		SUCCESS(退款成功)     */
	Status *string `json:"status,omitempty" required:"false" `
	/*
		        交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。
		fixed(一口价)
		auction(拍卖)
		guarantee_trade(一口价、拍卖)
		independent_simple_trade(旺店入门版交易)
		independent_shop_trade(旺店标准版交易)
		auto_delivery(自动发货)
		ec(直冲)
		cod(货到付款)
		fenxiao(分销)
		game_equipment(游戏装备)
		shopex_trade(ShopEX交易)
		netcn_trade(万网交易)
		external_trade(统一外部交易)     */
	Type *string `json:"type,omitempty" required:"false" `
	/*
	   页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 defalutValue��40    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoRefundsApplyGetRequest) SetFields(v []string) *TaobaoRefundsApplyGetRequest {
	s.Fields = &v
	return s
}
func (s *TaobaoRefundsApplyGetRequest) SetSellerNick(v string) *TaobaoRefundsApplyGetRequest {
	s.SellerNick = &v
	return s
}
func (s *TaobaoRefundsApplyGetRequest) SetStatus(v string) *TaobaoRefundsApplyGetRequest {
	s.Status = &v
	return s
}
func (s *TaobaoRefundsApplyGetRequest) SetType(v string) *TaobaoRefundsApplyGetRequest {
	s.Type = &v
	return s
}
func (s *TaobaoRefundsApplyGetRequest) SetPageNo(v int64) *TaobaoRefundsApplyGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoRefundsApplyGetRequest) SetPageSize(v int64) *TaobaoRefundsApplyGetRequest {
	s.PageSize = &v
	return s
}

func (req *TaobaoRefundsApplyGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	if req.SellerNick != nil {
		paramMap["seller_nick"] = *req.SellerNick
	}
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	if req.Type != nil {
		paramMap["type"] = *req.Type
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	return paramMap
}

func (req *TaobaoRefundsApplyGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
