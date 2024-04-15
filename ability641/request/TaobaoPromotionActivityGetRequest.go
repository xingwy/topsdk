package request


type TaobaoPromotionActivityGetRequest struct {
    /*
        活动的id     */
    ActivityId  *int64 `json:"activity_id,omitempty" required:"false" `
}

func (s *TaobaoPromotionActivityGetRequest) SetActivityId(v int64) *TaobaoPromotionActivityGetRequest {
    s.ActivityId = &v
    return s
}

func (req *TaobaoPromotionActivityGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActivityId != nil) {
        paramMap["activity_id"] = *req.ActivityId
    }
    return paramMap
}

func (req *TaobaoPromotionActivityGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}