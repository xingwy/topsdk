package domain


type TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo struct {
    /*
        审核不通过原因     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        短信签名状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        短信签名     */
    Signature  *string `json:"signature,omitempty" `

}

func (s *TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo) SetErrorMessage(v string) *TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo {
    s.ErrorMessage = &v
    return s
}
func (s *TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo) SetStatus(v int64) *TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo {
    s.Status = &v
    return s
}
func (s *TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo) SetSignature(v string) *TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo {
    s.Signature = &v
    return s
}
