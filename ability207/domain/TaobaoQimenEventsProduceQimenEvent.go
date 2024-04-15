package domain


type TaobaoQimenEventsProduceQimenEvent struct {
    /*
        奇门事件对象     */
    Event  *TaobaoQimenEventsProduceEvent `json:"event,omitempty" `

}

func (s *TaobaoQimenEventsProduceQimenEvent) SetEvent(v TaobaoQimenEventsProduceEvent) *TaobaoQimenEventsProduceQimenEvent {
    s.Event = &v
    return s
}
