package request


type TaobaoRefundNegotiatereturnRenderRequest struct {
    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoRefundNegotiatereturnRenderRequest) SetRefundId(v int64) *TaobaoRefundNegotiatereturnRenderRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoRefundNegotiatereturnRenderRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoRefundNegotiatereturnRenderRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}