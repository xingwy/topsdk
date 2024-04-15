package request


type TaobaoTradePostageUpdateRequest struct {
    /*
        主订单编号     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        邮费价格(邮费单位是元）     */
    PostFee  *string `json:"post_fee" required:"true" `
}

func (s *TaobaoTradePostageUpdateRequest) SetTid(v int64) *TaobaoTradePostageUpdateRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoTradePostageUpdateRequest) SetPostFee(v string) *TaobaoTradePostageUpdateRequest {
    s.PostFee = &v
    return s
}

func (req *TaobaoTradePostageUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.PostFee != nil) {
        paramMap["post_fee"] = *req.PostFee
    }
    return paramMap
}

func (req *TaobaoTradePostageUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}