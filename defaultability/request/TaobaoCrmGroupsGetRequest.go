package request


type TaobaoCrmGroupsGetRequest struct {
    /*
        每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1     */
    CurrentPage  *int64 `json:"current_page" required:"true" `
}

func (s *TaobaoCrmGroupsGetRequest) SetPageSize(v int64) *TaobaoCrmGroupsGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoCrmGroupsGetRequest) SetCurrentPage(v int64) *TaobaoCrmGroupsGetRequest {
    s.CurrentPage = &v
    return s
}

func (req *TaobaoCrmGroupsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    return paramMap
}

func (req *TaobaoCrmGroupsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}