package request


type AlibabaGpuSchemaCatsearchRequest struct {
    /*
        叶子类目ID     */
    LeafCatId  *int64 `json:"leaf_cat_id" required:"true" `
    /*
        当前页     */
    CurrentPage  *int64 `json:"current_page" required:"true" `
    /*
        每页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        渠道Id，如0代表天猫，8代表淘宝     */
    ProviderId  *int64 `json:"provider_id" required:"true" `
}

func (s *AlibabaGpuSchemaCatsearchRequest) SetLeafCatId(v int64) *AlibabaGpuSchemaCatsearchRequest {
    s.LeafCatId = &v
    return s
}
func (s *AlibabaGpuSchemaCatsearchRequest) SetCurrentPage(v int64) *AlibabaGpuSchemaCatsearchRequest {
    s.CurrentPage = &v
    return s
}
func (s *AlibabaGpuSchemaCatsearchRequest) SetPageSize(v int64) *AlibabaGpuSchemaCatsearchRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaGpuSchemaCatsearchRequest) SetProviderId(v int64) *AlibabaGpuSchemaCatsearchRequest {
    s.ProviderId = &v
    return s
}

func (req *AlibabaGpuSchemaCatsearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LeafCatId != nil) {
        paramMap["leaf_cat_id"] = *req.LeafCatId
    }
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.ProviderId != nil) {
        paramMap["provider_id"] = *req.ProviderId
    }
    return paramMap
}

func (req *AlibabaGpuSchemaCatsearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}