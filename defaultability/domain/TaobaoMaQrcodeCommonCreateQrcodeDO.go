package domain


type TaobaoMaQrcodeCommonCreateQrcodeDO struct {
    /*
        二维码所属渠道ID     */
    ChannelId  *int64 `json:"channel_id,omitempty" `

    /*
        二维码对应的渠道名     */
    ChannelName  *string `json:"channel_name,omitempty" `

    /*
        二维码的矢量图下载地址，只有设置入参need_eps为true且style不为官方模板时，才返回值     */
    EpsUrl  *string `json:"eps_url,omitempty" `

    /*
        二维码id     */
    QrcodeId  *int64 `json:"qrcode_id,omitempty" `

    /*
        二维码图片链接     */
    QrcodeUrl  *string `json:"qrcode_url,omitempty" `

    /*
        二维码的短地址，每一个原始地址都会生成一个短地址     */
    ShortUrl  *string `json:"short_url,omitempty" `

    /*
        二维码扫码后访问的目的地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetChannelId(v int64) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.ChannelId = &v
    return s
}
func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetChannelName(v string) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.ChannelName = &v
    return s
}
func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetEpsUrl(v string) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.EpsUrl = &v
    return s
}
func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetQrcodeId(v int64) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.QrcodeId = &v
    return s
}
func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetQrcodeUrl(v string) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.QrcodeUrl = &v
    return s
}
func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetShortUrl(v string) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.ShortUrl = &v
    return s
}
func (s *TaobaoMaQrcodeCommonCreateQrcodeDO) SetUrl(v string) *TaobaoMaQrcodeCommonCreateQrcodeDO {
    s.Url = &v
    return s
}
