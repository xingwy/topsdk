package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoRefundDetailGetRequest struct {
	/*
	   退款编号     */
	RefundId *int64 `json:"refund_id" required:"true" `
	/*
	   需要返回的字段。目前支持有：allowedOperations,refund_version     */
	Fields *[]string `json:"fields" required:"true" `
}

func (s *TaobaoRefundDetailGetRequest) SetRefundId(v int64) *TaobaoRefundDetailGetRequest {
	s.RefundId = &v
	return s
}
func (s *TaobaoRefundDetailGetRequest) SetFields(v []string) *TaobaoRefundDetailGetRequest {
	s.Fields = &v
	return s
}

func (req *TaobaoRefundDetailGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefundId != nil {
		paramMap["refund_id"] = *req.RefundId
	}
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	return paramMap
}

func (req *TaobaoRefundDetailGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
