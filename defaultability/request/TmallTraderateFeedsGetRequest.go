package request


type TmallTraderateFeedsGetRequest struct {
    /*
        交易子订单ID     */
    ChildTradeId  *int64 `json:"child_trade_id" required:"true" `
}

func (s *TmallTraderateFeedsGetRequest) SetChildTradeId(v int64) *TmallTraderateFeedsGetRequest {
    s.ChildTradeId = &v
    return s
}

func (req *TmallTraderateFeedsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ChildTradeId != nil) {
        paramMap["child_trade_id"] = *req.ChildTradeId
    }
    return paramMap
}

func (req *TmallTraderateFeedsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}