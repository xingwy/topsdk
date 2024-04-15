package request


type TmallItemCalculateHscodeGetRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
}

func (s *TmallItemCalculateHscodeGetRequest) SetItemId(v int64) *TmallItemCalculateHscodeGetRequest {
    s.ItemId = &v
    return s
}

func (req *TmallItemCalculateHscodeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TmallItemCalculateHscodeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}