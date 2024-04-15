package request


type AlibabaAliqinTaSmsNumQueryRequest struct {
    /*
        短信发送流水     */
    BizId  *string `json:"biz_id,omitempty" required:"false" `
    /*
        短信接收号码     */
    RecNum  *string `json:"rec_num" required:"true" `
    /*
        短信发送日期，支持近30天记录查询，格式yyyyMMdd     */
    QueryDate  *string `json:"query_date" required:"true" `
    /*
        分页参数,页码     */
    CurrentPage  *int64 `json:"current_page" required:"true" `
    /*
        分页参数，每页数量。最大值50     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAliqinTaSmsNumQueryRequest) SetBizId(v string) *AlibabaAliqinTaSmsNumQueryRequest {
    s.BizId = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumQueryRequest) SetRecNum(v string) *AlibabaAliqinTaSmsNumQueryRequest {
    s.RecNum = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumQueryRequest) SetQueryDate(v string) *AlibabaAliqinTaSmsNumQueryRequest {
    s.QueryDate = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumQueryRequest) SetCurrentPage(v int64) *AlibabaAliqinTaSmsNumQueryRequest {
    s.CurrentPage = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumQueryRequest) SetPageSize(v int64) *AlibabaAliqinTaSmsNumQueryRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAliqinTaSmsNumQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizId != nil) {
        paramMap["biz_id"] = *req.BizId
    }
    if(req.RecNum != nil) {
        paramMap["rec_num"] = *req.RecNum
    }
    if(req.QueryDate != nil) {
        paramMap["query_date"] = *req.QueryDate
    }
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAliqinTaSmsNumQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}