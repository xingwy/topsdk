package domain


type TaobaoRdsDbGetdbResultSet struct {
    /*
        results     */
    Results  *[]string `json:"results,omitempty" `

    /*
        totalResults     */
    TotalResults  *int64 `json:"total_results,omitempty" `

    /*
        exception     */
    Exception  *string `json:"exception,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoRdsDbGetdbResultSet) SetResults(v []string) *TaobaoRdsDbGetdbResultSet {
    s.Results = &v
    return s
}
func (s *TaobaoRdsDbGetdbResultSet) SetTotalResults(v int64) *TaobaoRdsDbGetdbResultSet {
    s.TotalResults = &v
    return s
}
func (s *TaobaoRdsDbGetdbResultSet) SetException(v string) *TaobaoRdsDbGetdbResultSet {
    s.Exception = &v
    return s
}
func (s *TaobaoRdsDbGetdbResultSet) SetErrorCode(v string) *TaobaoRdsDbGetdbResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRdsDbGetdbResultSet) SetErrorMsg(v string) *TaobaoRdsDbGetdbResultSet {
    s.ErrorMsg = &v
    return s
}
