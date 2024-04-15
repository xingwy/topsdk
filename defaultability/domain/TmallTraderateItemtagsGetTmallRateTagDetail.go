package domain


type TmallTraderateItemtagsGetTmallRateTagDetail struct {
    /*
        反应该标签的评价数量     */
    Count  *int64 `json:"count,omitempty" `

    /*
        标签的极性：1正极 -1负极     */
    Posi  *bool `json:"posi,omitempty" `

    /*
        标签名称     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TmallTraderateItemtagsGetTmallRateTagDetail) SetCount(v int64) *TmallTraderateItemtagsGetTmallRateTagDetail {
    s.Count = &v
    return s
}
func (s *TmallTraderateItemtagsGetTmallRateTagDetail) SetPosi(v bool) *TmallTraderateItemtagsGetTmallRateTagDetail {
    s.Posi = &v
    return s
}
func (s *TmallTraderateItemtagsGetTmallRateTagDetail) SetTagName(v string) *TmallTraderateItemtagsGetTmallRateTagDetail {
    s.TagName = &v
    return s
}
