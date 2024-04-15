package domain


type AlibabaAscpLogisticsIdentcodeUploadMptResult struct {
    /*
        请求状态码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        描述信息，当返回码不为0时，表示错误信息     */
    Msg  *string `json:"msg,omitempty" `

}

func (s *AlibabaAscpLogisticsIdentcodeUploadMptResult) SetCode(v int64) *AlibabaAscpLogisticsIdentcodeUploadMptResult {
    s.Code = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeUploadMptResult) SetMsg(v string) *AlibabaAscpLogisticsIdentcodeUploadMptResult {
    s.Msg = &v
    return s
}
