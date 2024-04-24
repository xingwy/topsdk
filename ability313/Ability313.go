package ability313

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability313/request"
	"github.com/xingwy/topsdk/ability313/response"
	"github.com/xingwy/topsdk/util"
)

type Ability313 struct {
	Client *topsdk.TopClient
}

func NewAbility313(client *topsdk.TopClient) *Ability313 {
	return &Ability313{client}
}

/*
根据用户nick获取开通的消息列表
*/
func (ability *Ability313) TaobaoJushitaJmsTopicsGet(req *request.TaobaoJushitaJmsTopicsGetRequest) (*response.TaobaoJushitaJmsTopicsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability313 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jushita.jms.topics.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJushitaJmsTopicsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询某个用户是否同步消息
*/
func (ability *Ability313) TaobaoJushitaJmsUserGet(req *request.TaobaoJushitaJmsUserGetRequest) (*response.TaobaoJushitaJmsUserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability313 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jushita.jms.user.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJushitaJmsUserGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
