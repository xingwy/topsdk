package request


type TmallTraderateItemtagsGetRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TmallTraderateItemtagsGetRequest) SetItemId(v int64) *TmallTraderateItemtagsGetRequest {
    s.ItemId = &v
    return s
}

func (req *TmallTraderateItemtagsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TmallTraderateItemtagsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}