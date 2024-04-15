package domain


type CainiaoCloudprintCustomareaUpdateKeyResult struct {
    /*
        keyName     */
    KeyName  *string `json:"key_name,omitempty" `

}

func (s *CainiaoCloudprintCustomareaUpdateKeyResult) SetKeyName(v string) *CainiaoCloudprintCustomareaUpdateKeyResult {
    s.KeyName = &v
    return s
}
