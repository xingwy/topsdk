package domain


type AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto struct {
    /*
        文件后缀名，文字内容转成txt，图片支持gif、jpg、png格式，音频支持mp3格式，视频支持mp4格式，目前只支持上传一个视频文件     */
    FileSuffix  *string `json:"file_suffix,omitempty" `

    /*
        文件名称，请带上后缀     */
    FileName  *string `json:"file_name,omitempty" `

    /*
        文件大小，单位：字节     */
    FileSize  *int64 `json:"file_size,omitempty" `

    /*
        文件二进制数组转base64，转的时候指定编码格式为UTF-8     */
    FileContents  *string `json:"file_contents,omitempty" `

}

func (s *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto) SetFileSuffix(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto {
    s.FileSuffix = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto) SetFileName(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto {
    s.FileName = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto) SetFileSize(v int64) *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto {
    s.FileSize = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto) SetFileContents(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto {
    s.FileContents = &v
    return s
}
