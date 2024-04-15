package domain


type TaobaoRefundMessagesGetPicUrl struct {
    /*
        图片链接地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoRefundMessagesGetPicUrl) SetUrl(v string) *TaobaoRefundMessagesGetPicUrl {
    s.Url = &v
    return s
}
