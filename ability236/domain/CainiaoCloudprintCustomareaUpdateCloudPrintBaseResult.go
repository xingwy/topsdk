package domain


type CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult struct {
    /*
        data     */
    Data  *CainiaoCloudprintCustomareaUpdateCustomAreaResult `json:"data,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMessage     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult) SetData(v CainiaoCloudprintCustomareaUpdateCustomAreaResult) *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult {
    s.Data = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult) SetErrorCode(v string) *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult) SetErrorMessage(v string) *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult {
    s.ErrorMessage = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult) SetSuccess(v bool) *CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult {
    s.Success = &v
    return s
}
