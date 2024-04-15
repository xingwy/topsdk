package request


type TaobaoItemCatpropsModificationGetRequest struct {
    /*
        类目Id（与商品Id二选一即可）     */
    CategoryId  *int64 `json:"category_id,omitempty" required:"false" `
    /*
        商品Id（与类目Id二选一即可。若同时传入商品Id和类目Id，则优先使用商品Id。若填写商品Id，则起始时间设为该商品最近修改时间）     */
    ItemId  *string `json:"item_id,omitempty" required:"false" `
    /*
        起始请求时间（建议传入，默认为90天内）     */
    StartTime  *string `json:"start_time,omitempty" required:"false" `
}

func (s *TaobaoItemCatpropsModificationGetRequest) SetCategoryId(v int64) *TaobaoItemCatpropsModificationGetRequest {
    s.CategoryId = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetRequest) SetItemId(v string) *TaobaoItemCatpropsModificationGetRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetRequest) SetStartTime(v string) *TaobaoItemCatpropsModificationGetRequest {
    s.StartTime = &v
    return s
}

func (req *TaobaoItemCatpropsModificationGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    return paramMap
}

func (req *TaobaoItemCatpropsModificationGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}