package domain


type TaobaoRdsDbCreateaccountResultSet struct {
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

func (s *TaobaoRdsDbCreateaccountResultSet) SetResults(v []string) *TaobaoRdsDbCreateaccountResultSet {
    s.Results = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountResultSet) SetTotalResults(v int64) *TaobaoRdsDbCreateaccountResultSet {
    s.TotalResults = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountResultSet) SetException(v string) *TaobaoRdsDbCreateaccountResultSet {
    s.Exception = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountResultSet) SetErrorCode(v string) *TaobaoRdsDbCreateaccountResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountResultSet) SetErrorMsg(v string) *TaobaoRdsDbCreateaccountResultSet {
    s.ErrorMsg = &v
    return s
}
