package request


type AlibabaGpuAddSchemaGetRequest struct {
    /*
        叶子类目ID     */
    LeafCatId  *int64 `json:"leaf_cat_id" required:"true" `
    /*
        品牌ID     */
    BrandId  *int64 `json:"brand_id,omitempty" required:"false" `
    /*
        当前用户所在渠道如0代表天猫，8代表淘宝     */
    ProviderId  *int64 `json:"provider_id" required:"true" `
}

func (s *AlibabaGpuAddSchemaGetRequest) SetLeafCatId(v int64) *AlibabaGpuAddSchemaGetRequest {
    s.LeafCatId = &v
    return s
}
func (s *AlibabaGpuAddSchemaGetRequest) SetBrandId(v int64) *AlibabaGpuAddSchemaGetRequest {
    s.BrandId = &v
    return s
}
func (s *AlibabaGpuAddSchemaGetRequest) SetProviderId(v int64) *AlibabaGpuAddSchemaGetRequest {
    s.ProviderId = &v
    return s
}

func (req *AlibabaGpuAddSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LeafCatId != nil) {
        paramMap["leaf_cat_id"] = *req.LeafCatId
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    if(req.ProviderId != nil) {
        paramMap["provider_id"] = *req.ProviderId
    }
    return paramMap
}

func (req *AlibabaGpuAddSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}