package domain


type TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO struct {
    /*
        文件名称     */
    FileName  *string `json:"file_name,omitempty" `

    /*
        文件大小     */
    FileSize  *int64 `json:"file_size,omitempty" `

    /*
        文件后缀名     */
    FileSuffix  *string `json:"file_suffix,omitempty" `

    /*
        文件Base64转码后的字符串     */
    FileContents  *string `json:"file_contents,omitempty" `

}

func (s *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO) SetFileName(v string) *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO {
    s.FileName = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO) SetFileSize(v int64) *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO {
    s.FileSize = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO) SetFileSuffix(v string) *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO {
    s.FileSuffix = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO) SetFileContents(v string) *TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO {
    s.FileContents = &v
    return s
}
