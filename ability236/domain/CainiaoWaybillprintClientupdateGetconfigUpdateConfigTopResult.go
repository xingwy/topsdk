package domain


type CainiaoWaybillprintClientupdateGetconfigUpdateConfigTopResult struct {
    /*
        status     */
    Status  *int64 `json:"status,omitempty" `

    /*
        description     */
    Description  *string `json:"description,omitempty" `

}

func (s *CainiaoWaybillprintClientupdateGetconfigUpdateConfigTopResult) SetStatus(v int64) *CainiaoWaybillprintClientupdateGetconfigUpdateConfigTopResult {
    s.Status = &v
    return s
}
func (s *CainiaoWaybillprintClientupdateGetconfigUpdateConfigTopResult) SetDescription(v string) *CainiaoWaybillprintClientupdateGetconfigUpdateConfigTopResult {
    s.Description = &v
    return s
}
