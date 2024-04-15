package request


type TmallItemCombineGetRequest struct {
    /*
        组合商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
}

func (s *TmallItemCombineGetRequest) SetItemId(v int64) *TmallItemCombineGetRequest {
    s.ItemId = &v
    return s
}

func (req *TmallItemCombineGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TmallItemCombineGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}