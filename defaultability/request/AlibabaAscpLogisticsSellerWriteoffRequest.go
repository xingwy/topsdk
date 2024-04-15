package request


type AlibabaAscpLogisticsSellerWriteoffRequest struct {
    /*
        所要核销的核销物流单ID，在alibaba.ascp.logistics.seller.orders.get中获取。     */
    LpOrderId  *int64 `json:"lp_order_id" required:"true" `
    /*
        核销码     */
    ReceiveCode  *string `json:"receive_code" required:"true" `
}

func (s *AlibabaAscpLogisticsSellerWriteoffRequest) SetLpOrderId(v int64) *AlibabaAscpLogisticsSellerWriteoffRequest {
    s.LpOrderId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWriteoffRequest) SetReceiveCode(v string) *AlibabaAscpLogisticsSellerWriteoffRequest {
    s.ReceiveCode = &v
    return s
}

func (req *AlibabaAscpLogisticsSellerWriteoffRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LpOrderId != nil) {
        paramMap["lp_order_id"] = *req.LpOrderId
    }
    if(req.ReceiveCode != nil) {
        paramMap["receive_code"] = *req.ReceiveCode
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsSellerWriteoffRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}