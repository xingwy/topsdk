package request


type TaobaoRefundNegotiatereturnRequest struct {
    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        退款版本号     */
    RefundVersion  *int64 `json:"refund_version" required:"true" `
    /*
        退款金额，单位（分）     */
    RefundFee  *int64 `json:"refund_fee" required:"true" `
    /*
        地址ID，通过协商退货退款渲染接口获取到的     */
    AddressId  *int64 `json:"address_id" required:"true" `
}

func (s *TaobaoRefundNegotiatereturnRequest) SetRefundId(v int64) *TaobaoRefundNegotiatereturnRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRequest) SetRefundVersion(v int64) *TaobaoRefundNegotiatereturnRequest {
    s.RefundVersion = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRequest) SetRefundFee(v int64) *TaobaoRefundNegotiatereturnRequest {
    s.RefundFee = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRequest) SetAddressId(v int64) *TaobaoRefundNegotiatereturnRequest {
    s.AddressId = &v
    return s
}

func (req *TaobaoRefundNegotiatereturnRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundVersion != nil) {
        paramMap["refund_version"] = *req.RefundVersion
    }
    if(req.RefundFee != nil) {
        paramMap["refund_fee"] = *req.RefundFee
    }
    if(req.AddressId != nil) {
        paramMap["address_id"] = *req.AddressId
    }
    return paramMap
}

func (req *TaobaoRefundNegotiatereturnRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}