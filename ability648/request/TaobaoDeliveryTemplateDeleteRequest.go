package request


type TaobaoDeliveryTemplateDeleteRequest struct {
    /*
        运费模板ID     */
    TemplateId  *int64 `json:"template_id" required:"true" `
}

func (s *TaobaoDeliveryTemplateDeleteRequest) SetTemplateId(v int64) *TaobaoDeliveryTemplateDeleteRequest {
    s.TemplateId = &v
    return s
}

func (req *TaobaoDeliveryTemplateDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TemplateId != nil) {
        paramMap["template_id"] = *req.TemplateId
    }
    return paramMap
}

func (req *TaobaoDeliveryTemplateDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}