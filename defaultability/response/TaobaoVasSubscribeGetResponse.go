package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoVasSubscribeGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   用户订购信息
	*/
	ArticleUserSubscribes []domain.TaobaoVasSubscribeGetArticleUserSubscribe `json:"article_user_subscribes,omitempty" `
}
