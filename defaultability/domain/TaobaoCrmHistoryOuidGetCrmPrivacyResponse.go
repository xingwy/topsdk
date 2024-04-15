package domain


type TaobaoCrmHistoryOuidGetCrmPrivacyResponse struct {
    /*
        ouid     */
    Ouid  *string `json:"ouid,omitempty" `

}

func (s *TaobaoCrmHistoryOuidGetCrmPrivacyResponse) SetOuid(v string) *TaobaoCrmHistoryOuidGetCrmPrivacyResponse {
    s.Ouid = &v
    return s
}
