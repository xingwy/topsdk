package request


type TaobaoTradeSellermemosGetRequest struct {
    /*
        主订单ID     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoTradeSellermemosGetRequest) SetTid(v int64) *TaobaoTradeSellermemosGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoTradeSellermemosGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoTradeSellermemosGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}