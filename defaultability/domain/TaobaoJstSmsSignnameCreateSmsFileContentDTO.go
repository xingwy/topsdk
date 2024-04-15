package domain


type TaobaoJstSmsSignnameCreateSmsFileContentDTO struct {
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

func (s *TaobaoJstSmsSignnameCreateSmsFileContentDTO) SetFileName(v string) *TaobaoJstSmsSignnameCreateSmsFileContentDTO {
    s.FileName = &v
    return s
}
func (s *TaobaoJstSmsSignnameCreateSmsFileContentDTO) SetFileSize(v int64) *TaobaoJstSmsSignnameCreateSmsFileContentDTO {
    s.FileSize = &v
    return s
}
func (s *TaobaoJstSmsSignnameCreateSmsFileContentDTO) SetFileSuffix(v string) *TaobaoJstSmsSignnameCreateSmsFileContentDTO {
    s.FileSuffix = &v
    return s
}
func (s *TaobaoJstSmsSignnameCreateSmsFileContentDTO) SetFileContents(v string) *TaobaoJstSmsSignnameCreateSmsFileContentDTO {
    s.FileContents = &v
    return s
}
