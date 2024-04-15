package request


type TaobaoCrmHistoryOuidGetRequest struct {
    /*
        买家淘宝Nick     */
    BuyerNick  *string `json:"buyer_nick" required:"true" `
}

func (s *TaobaoCrmHistoryOuidGetRequest) SetBuyerNick(v string) *TaobaoCrmHistoryOuidGetRequest {
    s.BuyerNick = &v
    return s
}

func (req *TaobaoCrmHistoryOuidGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    return paramMap
}

func (req *TaobaoCrmHistoryOuidGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}