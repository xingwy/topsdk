package domain


type TmallTraderateFeedsGetTags struct {
    /*
        表示标签的极性，正极true，负极false     */
    Posi  *bool `json:"posi,omitempty" `

    /*
        表示标签的名称     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TmallTraderateFeedsGetTags) SetPosi(v bool) *TmallTraderateFeedsGetTags {
    s.Posi = &v
    return s
}
func (s *TmallTraderateFeedsGetTags) SetTagName(v string) *TmallTraderateFeedsGetTags {
    s.TagName = &v
    return s
}
