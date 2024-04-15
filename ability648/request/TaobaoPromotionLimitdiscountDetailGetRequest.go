package request


type TaobaoPromotionLimitdiscountDetailGetRequest struct {
    /*
        限时打折ID。这个针对查询唯一限时打折情况。     */
    LimitDiscountId  *int64 `json:"limit_discount_id" required:"true" `
}

func (s *TaobaoPromotionLimitdiscountDetailGetRequest) SetLimitDiscountId(v int64) *TaobaoPromotionLimitdiscountDetailGetRequest {
    s.LimitDiscountId = &v
    return s
}

func (req *TaobaoPromotionLimitdiscountDetailGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LimitDiscountId != nil) {
        paramMap["limit_discount_id"] = *req.LimitDiscountId
    }
    return paramMap
}

func (req *TaobaoPromotionLimitdiscountDetailGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}