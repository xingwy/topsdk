package request


type AlibabaAscpLogisticsSellerOrdersGetRequest struct {
    /*
        核销日期     */
    WriteOffDate  *string `json:"write_off_date,omitempty" required:"false" `
    /*
        分页索引     */
    PageIndex  *int64 `json:"page_index,omitempty" required:"false" `
    /*
        收货码     */
    ReceiveCode  *string `json:"receive_code" required:"true" `
    /*
        分页大小     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        淘系交易id     */
    Tid  *string `json:"tid,omitempty" required:"false" `
    /*
        1代表未核销，2代表已核销     */
    WriteOffStatus  *string `json:"write_off_status,omitempty" required:"false" `
}

func (s *AlibabaAscpLogisticsSellerOrdersGetRequest) SetWriteOffDate(v string) *AlibabaAscpLogisticsSellerOrdersGetRequest {
    s.WriteOffDate = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetRequest) SetPageIndex(v int64) *AlibabaAscpLogisticsSellerOrdersGetRequest {
    s.PageIndex = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetRequest) SetReceiveCode(v string) *AlibabaAscpLogisticsSellerOrdersGetRequest {
    s.ReceiveCode = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetRequest) SetPageSize(v int64) *AlibabaAscpLogisticsSellerOrdersGetRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetRequest) SetTid(v string) *AlibabaAscpLogisticsSellerOrdersGetRequest {
    s.Tid = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetRequest) SetWriteOffStatus(v string) *AlibabaAscpLogisticsSellerOrdersGetRequest {
    s.WriteOffStatus = &v
    return s
}

func (req *AlibabaAscpLogisticsSellerOrdersGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.WriteOffDate != nil) {
        paramMap["write_off_date"] = *req.WriteOffDate
    }
    if(req.PageIndex != nil) {
        paramMap["page_index"] = *req.PageIndex
    }
    if(req.ReceiveCode != nil) {
        paramMap["receive_code"] = *req.ReceiveCode
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.WriteOffStatus != nil) {
        paramMap["write_off_status"] = *req.WriteOffStatus
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsSellerOrdersGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}