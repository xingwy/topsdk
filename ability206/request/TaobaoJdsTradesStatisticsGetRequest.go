package request


type TaobaoJdsTradesStatisticsGetRequest struct {
    /*
        查询的日期，格式如YYYYMMDD的日期对应的数字     */
    Date  *int64 `json:"date" required:"true" `
}

func (s *TaobaoJdsTradesStatisticsGetRequest) SetDate(v int64) *TaobaoJdsTradesStatisticsGetRequest {
    s.Date = &v
    return s
}

func (req *TaobaoJdsTradesStatisticsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Date != nil) {
        paramMap["date"] = *req.Date
    }
    return paramMap
}

func (req *TaobaoJdsTradesStatisticsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}