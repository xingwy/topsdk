package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTraderatesGetResults struct {
	/*
	   交易ID     */
	Tid *int64 `json:"tid,omitempty" `

	/*
	   子订单ID     */
	Oid *int64 `json:"oid,omitempty" `

	/*
	   评价者角色.可选值:seller(卖家),buyer(买家)     */
	Role *string `json:"role,omitempty" `

	/*
	   评价者昵称     */
	Nick *string `json:"nick,omitempty" `

	/*
	   评价结果,可选值:good(好评),neutral(中评),bad(差评)     */
	Result *string `json:"result,omitempty" `

	/*
	   评价创建时间,格式:yyyy-MM-dd HH:mm:ss     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   被评价者昵称     */
	RatedNick *string `json:"rated_nick,omitempty" `

	/*
	   商品标题     */
	ItemTitle *string `json:"item_title,omitempty" `

	/*
	   商品价格,精确到2位小数;单位:元.如:200.07，表示:200元7分     */
	ItemPrice *string `json:"item_price,omitempty" `

	/*
	   评价内容,最大长度:500个汉字     */
	Content *string `json:"content,omitempty" `

	/*
	   评价解释,最大长度:500个汉字     */
	Reply *string `json:"reply,omitempty" `

	/*
	        评价信息是否用于记分，
	可取值：true(参与记分)和false(不参与记分)     */
	ValidScore *bool `json:"valid_score,omitempty" `

	/*
	   商品的数字ID     */
	NumIid *int64 `json:"num_iid,omitempty" `

	/*
	   物流服务评分     */
	LogisticsServiceScore *int64 `json:"logistics_service_score,omitempty" `

	/*
	   ouid     */
	Ouid *string `json:"ouid,omitempty" `
}

func (s *TaobaoTraderatesGetResults) SetTid(v int64) *TaobaoTraderatesGetResults {
	s.Tid = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetOid(v int64) *TaobaoTraderatesGetResults {
	s.Oid = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetRole(v string) *TaobaoTraderatesGetResults {
	s.Role = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetNick(v string) *TaobaoTraderatesGetResults {
	s.Nick = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetResult(v string) *TaobaoTraderatesGetResults {
	s.Result = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetCreated(v util.LocalTime) *TaobaoTraderatesGetResults {
	s.Created = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetRatedNick(v string) *TaobaoTraderatesGetResults {
	s.RatedNick = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetItemTitle(v string) *TaobaoTraderatesGetResults {
	s.ItemTitle = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetItemPrice(v string) *TaobaoTraderatesGetResults {
	s.ItemPrice = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetContent(v string) *TaobaoTraderatesGetResults {
	s.Content = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetReply(v string) *TaobaoTraderatesGetResults {
	s.Reply = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetValidScore(v bool) *TaobaoTraderatesGetResults {
	s.ValidScore = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetNumIid(v int64) *TaobaoTraderatesGetResults {
	s.NumIid = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetLogisticsServiceScore(v int64) *TaobaoTraderatesGetResults {
	s.LogisticsServiceScore = &v
	return s
}
func (s *TaobaoTraderatesGetResults) SetOuid(v string) *TaobaoTraderatesGetResults {
	s.Ouid = &v
	return s
}
