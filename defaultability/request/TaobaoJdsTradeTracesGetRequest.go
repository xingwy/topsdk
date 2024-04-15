package request


type TaobaoJdsTradeTracesGetRequest struct {
    /*
        是否返回用户的状态(是否存在)     */
    ReturnUserStatus  *bool `json:"return_user_status,omitempty" required:"false" `
    /*
        淘宝的订单tid     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoJdsTradeTracesGetRequest) SetReturnUserStatus(v bool) *TaobaoJdsTradeTracesGetRequest {
    s.ReturnUserStatus = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetRequest) SetTid(v int64) *TaobaoJdsTradeTracesGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoJdsTradeTracesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ReturnUserStatus != nil) {
        paramMap["return_user_status"] = *req.ReturnUserStatus
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoJdsTradeTracesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}