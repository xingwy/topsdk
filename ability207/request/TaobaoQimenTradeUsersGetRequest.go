package request


type TaobaoQimenTradeUsersGetRequest struct {
    /*
        页数     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        每页的数量     */
    PageIndex  *int64 `json:"page_index" required:"true" `
}

func (s *TaobaoQimenTradeUsersGetRequest) SetPageSize(v int64) *TaobaoQimenTradeUsersGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoQimenTradeUsersGetRequest) SetPageIndex(v int64) *TaobaoQimenTradeUsersGetRequest {
    s.PageIndex = &v
    return s
}

func (req *TaobaoQimenTradeUsersGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageIndex != nil) {
        paramMap["page_index"] = *req.PageIndex
    }
    return paramMap
}

func (req *TaobaoQimenTradeUsersGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}