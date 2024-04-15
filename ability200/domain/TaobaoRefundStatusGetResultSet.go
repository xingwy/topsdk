package domain


type TaobaoRefundStatusGetResultSet struct {
    /*
        数组对象     */
    ResultList  *[]TaobaoRefundStatusGetQueryRefundStatusResponse `json:"result_list,omitempty" `

    /*
        错误码，没有表示无异常     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoRefundStatusGetResultSet) SetResultList(v []TaobaoRefundStatusGetQueryRefundStatusResponse) *TaobaoRefundStatusGetResultSet {
    s.ResultList = &v
    return s
}
func (s *TaobaoRefundStatusGetResultSet) SetErrorCode(v string) *TaobaoRefundStatusGetResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRefundStatusGetResultSet) SetErrorMsg(v string) *TaobaoRefundStatusGetResultSet {
    s.ErrorMsg = &v
    return s
}
