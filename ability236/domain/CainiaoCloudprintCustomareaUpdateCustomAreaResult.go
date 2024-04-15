package domain


type CainiaoCloudprintCustomareaUpdateCustomAreaResult struct {
    /*
        customAreaId     */
    CustomAreaId  *int64 `json:"custom_area_id,omitempty" `

    /*
        customAreaUrl     */
    CustomAreaUrl  *string `json:"custom_area_url,omitempty" `

    /*
        keys     */
    Keys  *[]CainiaoCloudprintCustomareaUpdateKeyResult `json:"keys,omitempty" `

}

func (s *CainiaoCloudprintCustomareaUpdateCustomAreaResult) SetCustomAreaId(v int64) *CainiaoCloudprintCustomareaUpdateCustomAreaResult {
    s.CustomAreaId = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateCustomAreaResult) SetCustomAreaUrl(v string) *CainiaoCloudprintCustomareaUpdateCustomAreaResult {
    s.CustomAreaUrl = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateCustomAreaResult) SetKeys(v []CainiaoCloudprintCustomareaUpdateKeyResult) *CainiaoCloudprintCustomareaUpdateCustomAreaResult {
    s.Keys = &v
    return s
}
