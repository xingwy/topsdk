package request

import (
        "topsdk/util"
    )

type TaobaoRefundMessagesGetRequest struct {
    /*
        需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        页码 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        退款单号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。     */
    RefundPhase  *string `json:"refund_phase,omitempty" required:"false" `
}

func (s *TaobaoRefundMessagesGetRequest) SetFields(v []string) *TaobaoRefundMessagesGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoRefundMessagesGetRequest) SetPageNo(v int64) *TaobaoRefundMessagesGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoRefundMessagesGetRequest) SetPageSize(v int64) *TaobaoRefundMessagesGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoRefundMessagesGetRequest) SetRefundId(v int64) *TaobaoRefundMessagesGetRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRefundMessagesGetRequest) SetRefundPhase(v string) *TaobaoRefundMessagesGetRequest {
    s.RefundPhase = &v
    return s
}

func (req *TaobaoRefundMessagesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundPhase != nil) {
        paramMap["refund_phase"] = *req.RefundPhase
    }
    return paramMap
}

func (req *TaobaoRefundMessagesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}