package domain


type TaobaoCrmHistoryOmidGetCrmPrivacyResponse struct {
    /*
        omid     */
    Omid  *string `json:"omid,omitempty" `

}

func (s *TaobaoCrmHistoryOmidGetCrmPrivacyResponse) SetOmid(v string) *TaobaoCrmHistoryOmidGetCrmPrivacyResponse {
    s.Omid = &v
    return s
}
