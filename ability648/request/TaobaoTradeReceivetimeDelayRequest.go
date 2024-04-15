package request


type TaobaoTradeReceivetimeDelayRequest struct {
    /*
        主订单号     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        延长收货的天数，可选值为：3, 5, 7, 10。     */
    Days  *int64 `json:"days" required:"true" `
}

func (s *TaobaoTradeReceivetimeDelayRequest) SetTid(v int64) *TaobaoTradeReceivetimeDelayRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeReceivetimeDelayRequest) SetDays(v int64) *TaobaoTradeReceivetimeDelayRequest {
    s.Days = &v
    return s
}

func (req *TaobaoTradeReceivetimeDelayRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.Days != nil) {
        paramMap["days"] = *req.Days
    }
    return paramMap
}

func (req *TaobaoTradeReceivetimeDelayRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}