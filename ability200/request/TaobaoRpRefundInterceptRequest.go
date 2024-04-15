package request


type TaobaoRpRefundInterceptRequest struct {
    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        退款版本号     */
    RefundVersion  *int64 `json:"refund_version" required:"true" `
}

func (s *TaobaoRpRefundInterceptRequest) SetRefundId(v int64) *TaobaoRpRefundInterceptRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRpRefundInterceptRequest) SetRefundVersion(v int64) *TaobaoRpRefundInterceptRequest {
    s.RefundVersion = &v
    return s
}

func (req *TaobaoRpRefundInterceptRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundVersion != nil) {
        paramMap["refund_version"] = *req.RefundVersion
    }
    return paramMap
}

func (req *TaobaoRpRefundInterceptRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}