package request

import (
        "topsdk/util"
    )

type TaobaoFuwuScoresGetRequest struct {
    /*
        评价日期，查询某一天的评价     */
    Date  *util.LocalTime `json:"date" required:"true" `
    /*
        每页获取条数。默认值40，最小值1，最大值100。 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        当前页     */
    CurrentPage  *int64 `json:"current_page" required:"true" `
}

func (s *TaobaoFuwuScoresGetRequest) SetDate(v util.LocalTime) *TaobaoFuwuScoresGetRequest {
    s.Date = &v
    return s
}
func (s *TaobaoFuwuScoresGetRequest) SetPageSize(v int64) *TaobaoFuwuScoresGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFuwuScoresGetRequest) SetCurrentPage(v int64) *TaobaoFuwuScoresGetRequest {
    s.CurrentPage = &v
    return s
}

func (req *TaobaoFuwuScoresGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Date != nil) {
        paramMap["date"] = *req.Date
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    return paramMap
}

func (req *TaobaoFuwuScoresGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}