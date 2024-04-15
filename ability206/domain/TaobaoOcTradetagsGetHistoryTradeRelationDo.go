package domain

import (
        "topsdk/util"
    )

type TaobaoOcTradetagsGetHistoryTradeRelationDo struct {
    /*
        订单标签记录id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        订单id     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        标签名称     */
    TagName  *string `json:"tag_name,omitempty" `

    /*
        标签类型       1：官方标签      2：自定义标签     */
    TagType  *int64 `json:"tag_type,omitempty" `

    /*
        标签值，json格式     */
    TagValue  *string `json:"tag_value,omitempty" `

    /*
        记录的创建时间     */
    GmtCreated  *util.LocalTime `json:"gmt_created,omitempty" `

    /*
        记录的最新修改时间     */
    GmtModified  *util.LocalTime `json:"gmt_modified,omitempty" `

    /*
        该标签在消费者端是否显示,0:不显示,1：显示     */
    Visible  *int64 `json:"visible,omitempty" `

}

func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetId(v int64) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.Id = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetTid(v int64) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.Tid = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetTagName(v string) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.TagName = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetTagType(v int64) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.TagType = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetTagValue(v string) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.TagValue = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetGmtCreated(v util.LocalTime) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.GmtCreated = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetGmtModified(v util.LocalTime) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.GmtModified = &v
    return s
}
func (s *TaobaoOcTradetagsGetHistoryTradeRelationDo) SetVisible(v int64) *TaobaoOcTradetagsGetHistoryTradeRelationDo {
    s.Visible = &v
    return s
}
