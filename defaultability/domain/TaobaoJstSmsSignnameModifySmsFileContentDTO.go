package domain


type TaobaoJstSmsSignnameModifySmsFileContentDTO struct {
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

func (s *TaobaoJstSmsSignnameModifySmsFileContentDTO) SetFileName(v string) *TaobaoJstSmsSignnameModifySmsFileContentDTO {
    s.FileName = &v
    return s
}
func (s *TaobaoJstSmsSignnameModifySmsFileContentDTO) SetFileSize(v int64) *TaobaoJstSmsSignnameModifySmsFileContentDTO {
    s.FileSize = &v
    return s
}
func (s *TaobaoJstSmsSignnameModifySmsFileContentDTO) SetFileSuffix(v string) *TaobaoJstSmsSignnameModifySmsFileContentDTO {
    s.FileSuffix = &v
    return s
}
func (s *TaobaoJstSmsSignnameModifySmsFileContentDTO) SetFileContents(v string) *TaobaoJstSmsSignnameModifySmsFileContentDTO {
    s.FileContents = &v
    return s
}
