package request


type AlibabaGpuSchemaAddRequest struct {
    /*
        叶子类目ID     */
    LeafCatId  *int64 `json:"leaf_cat_id" required:"true" `
    /*
        品牌ID     */
    BrandId  *int64 `json:"brand_id,omitempty" required:"false" `
    /*
        根据alibaba.gpu.add.schema.get获取的规则提交上来的schema     */
    SchemaXmlFields  *string `json:"schema_xml_fields" required:"true" `
    /*
        当前用户所在渠道如0代表天猫，8代表淘宝     */
    ProviderId  *int64 `json:"provider_id" required:"true" `
}

func (s *AlibabaGpuSchemaAddRequest) SetLeafCatId(v int64) *AlibabaGpuSchemaAddRequest {
    s.LeafCatId = &v
    return s
}
func (s *AlibabaGpuSchemaAddRequest) SetBrandId(v int64) *AlibabaGpuSchemaAddRequest {
    s.BrandId = &v
    return s
}
func (s *AlibabaGpuSchemaAddRequest) SetSchemaXmlFields(v string) *AlibabaGpuSchemaAddRequest {
    s.SchemaXmlFields = &v
    return s
}
func (s *AlibabaGpuSchemaAddRequest) SetProviderId(v int64) *AlibabaGpuSchemaAddRequest {
    s.ProviderId = &v
    return s
}

func (req *AlibabaGpuSchemaAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LeafCatId != nil) {
        paramMap["leaf_cat_id"] = *req.LeafCatId
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    if(req.SchemaXmlFields != nil) {
        paramMap["schema_xml_fields"] = *req.SchemaXmlFields
    }
    if(req.ProviderId != nil) {
        paramMap["provider_id"] = *req.ProviderId
    }
    return paramMap
}

func (req *AlibabaGpuSchemaAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}