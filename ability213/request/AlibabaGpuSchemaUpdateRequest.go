package request


type AlibabaGpuSchemaUpdateRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        更新产品提交的schema数据     */
    SchemaXmlFields  *string `json:"schema_xml_fields" required:"true" `
    /*
        当前用户所在渠道如0代表天猫，8代表淘宝     */
    ProviderId  *int64 `json:"provider_id" required:"true" `
}

func (s *AlibabaGpuSchemaUpdateRequest) SetProductId(v int64) *AlibabaGpuSchemaUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *AlibabaGpuSchemaUpdateRequest) SetSchemaXmlFields(v string) *AlibabaGpuSchemaUpdateRequest {
    s.SchemaXmlFields = &v
    return s
}
func (s *AlibabaGpuSchemaUpdateRequest) SetProviderId(v int64) *AlibabaGpuSchemaUpdateRequest {
    s.ProviderId = &v
    return s
}

func (req *AlibabaGpuSchemaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.SchemaXmlFields != nil) {
        paramMap["schema_xml_fields"] = *req.SchemaXmlFields
    }
    if(req.ProviderId != nil) {
        paramMap["provider_id"] = *req.ProviderId
    }
    return paramMap
}

func (req *AlibabaGpuSchemaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}