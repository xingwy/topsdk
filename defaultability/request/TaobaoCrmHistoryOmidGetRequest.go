package request


type TaobaoCrmHistoryOmidGetRequest struct {
    /*
        买家淘宝Nick     */
    BuyerNick  *string `json:"buyer_nick" required:"true" `
}

func (s *TaobaoCrmHistoryOmidGetRequest) SetBuyerNick(v string) *TaobaoCrmHistoryOmidGetRequest {
    s.BuyerNick = &v
    return s
}

func (req *TaobaoCrmHistoryOmidGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    return paramMap
}

func (req *TaobaoCrmHistoryOmidGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}