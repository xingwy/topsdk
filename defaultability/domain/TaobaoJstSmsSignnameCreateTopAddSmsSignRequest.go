package domain


type TaobaoJstSmsSignnameCreateTopAddSmsSignRequest struct {
    /*
        签名发送自带【】符号，无须添加【】、()、[]符号，避免重复 不支持如“客户服务”、“友情提醒”等过于宽泛内容、不支持“测试”字样的签名     */
    SignName  *string `json:"sign_name,omitempty" `

    /*
        请输入签名用途（必填）、企业官网链接（可提升通过率）     */
    Remark  *string `json:"remark,omitempty" `

    /*
        0--企事业单位的全程或简称   1--已备案网站  2--已上线APP  3--公众号或小程序 4--电商平台店铺名 5--已注册商标名     */
    SignSource  *int64 `json:"sign_source,omitempty" `

    /*
        上传的证明文件     */
    SignFileList  *[]TaobaoJstSmsSignnameCreateSmsFileContentDTO `json:"sign_file_list,omitempty" `

}

func (s *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest) SetSignName(v string) *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest {
    s.SignName = &v
    return s
}
func (s *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest) SetRemark(v string) *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest) SetSignSource(v int64) *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest {
    s.SignSource = &v
    return s
}
func (s *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest) SetSignFileList(v []TaobaoJstSmsSignnameCreateSmsFileContentDTO) *TaobaoJstSmsSignnameCreateTopAddSmsSignRequest {
    s.SignFileList = &v
    return s
}
