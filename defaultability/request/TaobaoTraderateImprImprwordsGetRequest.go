package request


type TaobaoTraderateImprImprwordsGetRequest struct {
    /*
        淘宝一级类目id     */
    CatRootId  *int64 `json:"cat_root_id" required:"true" `
    /*
        淘宝叶子类目id     */
    CatLeafId  *int64 `json:"cat_leaf_id,omitempty" required:"false" `
}

func (s *TaobaoTraderateImprImprwordsGetRequest) SetCatRootId(v int64) *TaobaoTraderateImprImprwordsGetRequest {
    s.CatRootId = &v
    return s
}
func (s *TaobaoTraderateImprImprwordsGetRequest) SetCatLeafId(v int64) *TaobaoTraderateImprImprwordsGetRequest {
    s.CatLeafId = &v
    return s
}

func (req *TaobaoTraderateImprImprwordsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CatRootId != nil) {
        paramMap["cat_root_id"] = *req.CatRootId
    }
    if(req.CatLeafId != nil) {
        paramMap["cat_leaf_id"] = *req.CatLeafId
    }
    return paramMap
}

func (req *TaobaoTraderateImprImprwordsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}