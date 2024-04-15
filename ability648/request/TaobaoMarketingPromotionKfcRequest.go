package request


type TaobaoMarketingPromotionKfcRequest struct {
    /*
        活动名称     */
    PromotionTitle  *string `json:"promotion_title" required:"true" `
    /*
        活动描述     */
    PromotionDesc  *string `json:"promotion_desc" required:"true" `
}

func (s *TaobaoMarketingPromotionKfcRequest) SetPromotionTitle(v string) *TaobaoMarketingPromotionKfcRequest {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoMarketingPromotionKfcRequest) SetPromotionDesc(v string) *TaobaoMarketingPromotionKfcRequest {
    s.PromotionDesc = &v
    return s
}

func (req *TaobaoMarketingPromotionKfcRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PromotionTitle != nil) {
        paramMap["promotion_title"] = *req.PromotionTitle
    }
    if(req.PromotionDesc != nil) {
        paramMap["promotion_desc"] = *req.PromotionDesc
    }
    return paramMap
}

func (req *TaobaoMarketingPromotionKfcRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}