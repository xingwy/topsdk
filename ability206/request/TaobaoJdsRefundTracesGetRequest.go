package request


type TaobaoJdsRefundTracesGetRequest struct {
    /*
        淘宝的退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        是否返回用户状态(是否存在)     */
    ReturnUserStatus  *bool `json:"return_user_status,omitempty" required:"false" `
}

func (s *TaobaoJdsRefundTracesGetRequest) SetRefundId(v int64) *TaobaoJdsRefundTracesGetRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoJdsRefundTracesGetRequest) SetReturnUserStatus(v bool) *TaobaoJdsRefundTracesGetRequest {
    s.ReturnUserStatus = &v
    return s
}

func (req *TaobaoJdsRefundTracesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.ReturnUserStatus != nil) {
        paramMap["return_user_status"] = *req.ReturnUserStatus
    }
    return paramMap
}

func (req *TaobaoJdsRefundTracesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}