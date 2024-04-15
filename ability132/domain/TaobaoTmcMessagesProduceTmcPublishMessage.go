package domain


type TaobaoTmcMessagesProduceTmcPublishMessage struct {
    /*
        消息内容的JSON表述，必须按照topic的定义来填充     */
    Content  *string `json:"content,omitempty" `

    /*
        消息的扩增属性，用json格式表示     */
    JsonExContent  *string `json:"json_ex_content,omitempty" `

    /*
        直发消息需要传入目标appkey     */
    TargetAppKey  *string `json:"target_app_key,omitempty" `

    /*
        目标分组     */
    TargetGroup  *string `json:"target_group,omitempty" `

    /*
        消息类型     */
    Topic  *string `json:"topic,omitempty" `

}

func (s *TaobaoTmcMessagesProduceTmcPublishMessage) SetContent(v string) *TaobaoTmcMessagesProduceTmcPublishMessage {
    s.Content = &v
    return s
}
func (s *TaobaoTmcMessagesProduceTmcPublishMessage) SetJsonExContent(v string) *TaobaoTmcMessagesProduceTmcPublishMessage {
    s.JsonExContent = &v
    return s
}
func (s *TaobaoTmcMessagesProduceTmcPublishMessage) SetTargetAppKey(v string) *TaobaoTmcMessagesProduceTmcPublishMessage {
    s.TargetAppKey = &v
    return s
}
func (s *TaobaoTmcMessagesProduceTmcPublishMessage) SetTargetGroup(v string) *TaobaoTmcMessagesProduceTmcPublishMessage {
    s.TargetGroup = &v
    return s
}
func (s *TaobaoTmcMessagesProduceTmcPublishMessage) SetTopic(v string) *TaobaoTmcMessagesProduceTmcPublishMessage {
    s.Topic = &v
    return s
}
