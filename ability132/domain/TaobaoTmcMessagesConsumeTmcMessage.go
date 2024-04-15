package domain

import (
        "topsdk/util"
    )

type TaobaoTmcMessagesConsumeTmcMessage struct {
    /*
        消息所属的用户编号     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        用户的昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        消息详细内容，格式为JSON/XML     */
    Content  *string `json:"content,omitempty" `

    /*
        消息ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        消息发布时间     */
    PubTime  *util.LocalTime `json:"pub_time,omitempty" `

    /*
        消息发布者的AppKey     */
    PubAppKey  *string `json:"pub_app_key,omitempty" `

    /*
        消息所属主题     */
    Topic  *string `json:"topic,omitempty" `

}

func (s *TaobaoTmcMessagesConsumeTmcMessage) SetUserId(v int64) *TaobaoTmcMessagesConsumeTmcMessage {
    s.UserId = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeTmcMessage) SetUserNick(v string) *TaobaoTmcMessagesConsumeTmcMessage {
    s.UserNick = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeTmcMessage) SetContent(v string) *TaobaoTmcMessagesConsumeTmcMessage {
    s.Content = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeTmcMessage) SetId(v int64) *TaobaoTmcMessagesConsumeTmcMessage {
    s.Id = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeTmcMessage) SetPubTime(v util.LocalTime) *TaobaoTmcMessagesConsumeTmcMessage {
    s.PubTime = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeTmcMessage) SetPubAppKey(v string) *TaobaoTmcMessagesConsumeTmcMessage {
    s.PubAppKey = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeTmcMessage) SetTopic(v string) *TaobaoTmcMessagesConsumeTmcMessage {
    s.Topic = &v
    return s
}
