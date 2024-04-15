package request


type CainiaoCloudprintCustomareaUpdateRequest struct {
    /*
        自定义区id（不可修改）     */
    CustomAreaId  *int64 `json:"custom_area_id" required:"true" `
    /*
        自定义区名称（可修改）     */
    CustomAreaName  *string `json:"custom_area_name" required:"true" `
    /*
        自定义区内容（可修改）     */
    CustomAreaContent  *string `json:"custom_area_content" required:"true" `
}

func (s *CainiaoCloudprintCustomareaUpdateRequest) SetCustomAreaId(v int64) *CainiaoCloudprintCustomareaUpdateRequest {
    s.CustomAreaId = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateRequest) SetCustomAreaName(v string) *CainiaoCloudprintCustomareaUpdateRequest {
    s.CustomAreaName = &v
    return s
}
func (s *CainiaoCloudprintCustomareaUpdateRequest) SetCustomAreaContent(v string) *CainiaoCloudprintCustomareaUpdateRequest {
    s.CustomAreaContent = &v
    return s
}

func (req *CainiaoCloudprintCustomareaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CustomAreaId != nil) {
        paramMap["custom_area_id"] = *req.CustomAreaId
    }
    if(req.CustomAreaName != nil) {
        paramMap["custom_area_name"] = *req.CustomAreaName
    }
    if(req.CustomAreaContent != nil) {
        paramMap["custom_area_content"] = *req.CustomAreaContent
    }
    return paramMap
}

func (req *CainiaoCloudprintCustomareaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}