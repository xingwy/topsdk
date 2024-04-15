package request


type TaobaoJdsTradesStatisticsDiffRequest struct {
    /*
        查询的日期，格式如YYYYMMDD的日期对应的数字     */
    Date  *int64 `json:"date" required:"true" `
    /*
        页数 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        需要比较的状态     */
    PostStatus  *string `json:"post_status" required:"true" `
    /*
        需要比较的状态，将会和post_status做比较     */
    PreStatus  *string `json:"pre_status" required:"true" `
}

func (s *TaobaoJdsTradesStatisticsDiffRequest) SetDate(v int64) *TaobaoJdsTradesStatisticsDiffRequest {
    s.Date = &v
    return s
}
func (s *TaobaoJdsTradesStatisticsDiffRequest) SetPageNo(v int64) *TaobaoJdsTradesStatisticsDiffRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoJdsTradesStatisticsDiffRequest) SetPostStatus(v string) *TaobaoJdsTradesStatisticsDiffRequest {
    s.PostStatus = &v
    return s
}
func (s *TaobaoJdsTradesStatisticsDiffRequest) SetPreStatus(v string) *TaobaoJdsTradesStatisticsDiffRequest {
    s.PreStatus = &v
    return s
}

func (req *TaobaoJdsTradesStatisticsDiffRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Date != nil) {
        paramMap["date"] = *req.Date
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PostStatus != nil) {
        paramMap["post_status"] = *req.PostStatus
    }
    if(req.PreStatus != nil) {
        paramMap["pre_status"] = *req.PreStatus
    }
    return paramMap
}

func (req *TaobaoJdsTradesStatisticsDiffRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}