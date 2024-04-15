package request


type AlibabaGpuUpdateSchemaGetRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        当前用户所在渠道如0代表天猫，8代表淘宝     */
    ProviderId  *int64 `json:"provider_id" required:"true" `
}

func (s *AlibabaGpuUpdateSchemaGetRequest) SetProductId(v int64) *AlibabaGpuUpdateSchemaGetRequest {
    s.ProductId = &v
    return s
}
func (s *AlibabaGpuUpdateSchemaGetRequest) SetProviderId(v int64) *AlibabaGpuUpdateSchemaGetRequest {
    s.ProviderId = &v
    return s
}

func (req *AlibabaGpuUpdateSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.ProviderId != nil) {
        paramMap["provider_id"] = *req.ProviderId
    }
    return paramMap
}

func (req *AlibabaGpuUpdateSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}