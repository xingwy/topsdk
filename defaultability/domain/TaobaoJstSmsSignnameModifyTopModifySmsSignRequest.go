package domain


type TaobaoJstSmsSignnameModifyTopModifySmsSignRequest struct {
    /*
        要修改的签名,不能修改签名     */
    SignName  *string `json:"sign_name,omitempty" `

    /*
        请输入签名用途（必填）、企业官网链接（可提升通过率）     */
    Remark  *string `json:"remark,omitempty" `

    /*
        0--企事业单位的全程或简称   1--已备案网站  2--已上线APP  3--公众号或小程序 4--电商平台店铺名 5--已注册商标名     */
    SignSource  *int64 `json:"sign_source,omitempty" `

    /*
        上传的证明文件     */
    SignFileList  *[]TaobaoJstSmsSignnameModifySmsFileContentDTO `json:"sign_file_list,omitempty" `

}

func (s *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest) SetSignName(v string) *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest {
    s.SignName = &v
    return s
}
func (s *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest) SetRemark(v string) *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest) SetSignSource(v int64) *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest {
    s.SignSource = &v
    return s
}
func (s *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest) SetSignFileList(v []TaobaoJstSmsSignnameModifySmsFileContentDTO) *TaobaoJstSmsSignnameModifyTopModifySmsSignRequest {
    s.SignFileList = &v
    return s
}
