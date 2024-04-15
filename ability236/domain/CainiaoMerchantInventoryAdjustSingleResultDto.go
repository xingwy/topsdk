package domain


type CainiaoMerchantInventoryAdjustSingleResultDto struct {
    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMessage     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        success     */
    Flag  *bool `json:"flag,omitempty" `

}

func (s *CainiaoMerchantInventoryAdjustSingleResultDto) SetErrorCode(v string) *CainiaoMerchantInventoryAdjustSingleResultDto {
    s.ErrorCode = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustSingleResultDto) SetErrorMessage(v string) *CainiaoMerchantInventoryAdjustSingleResultDto {
    s.ErrorMessage = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustSingleResultDto) SetFlag(v bool) *CainiaoMerchantInventoryAdjustSingleResultDto {
    s.Flag = &v
    return s
}
