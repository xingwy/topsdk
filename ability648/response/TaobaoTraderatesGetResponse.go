package response

import (
	"github.com/xingwy/topsdk/ability648/domain"
)

type TaobaoTraderatesGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据
	*/
	TotalResults int64 `json:"total_results,omitempty" `
	/*
	   当使用use_has_next时返回信息，如果还有下一页则返回true
	*/
	HasNext bool `json:"has_next,omitempty" `
	/*
	   评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息
	*/
	TradeRates []domain.TaobaoTraderatesGetResults `json:"trade_rates,omitempty" `
}
