package request


type TaobaoQimenTradeUserAddRequest struct {
    /*
        商家备注     */
    Memo  *string `json:"memo,omitempty" required:"false" `
}

func (s *TaobaoQimenTradeUserAddRequest) SetMemo(v string) *TaobaoQimenTradeUserAddRequest {
    s.Memo = &v
    return s
}

func (req *TaobaoQimenTradeUserAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Memo != nil) {
        paramMap["memo"] = *req.Memo
    }
    return paramMap
}

func (req *TaobaoQimenTradeUserAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}