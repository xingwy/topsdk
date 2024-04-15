package request


type TmallItemHscodeAuditResultsQueryRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TmallItemHscodeAuditResultsQueryRequest) SetItemId(v int64) *TmallItemHscodeAuditResultsQueryRequest {
    s.ItemId = &v
    return s
}

func (req *TmallItemHscodeAuditResultsQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TmallItemHscodeAuditResultsQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}