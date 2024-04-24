package ability207

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability207/request"
	"github.com/xingwy/topsdk/ability207/response"
	"github.com/xingwy/topsdk/util"
)

type Ability207 struct {
	Client *topsdk.TopClient
}

func NewAbility207(client *topsdk.TopClient) *Ability207 {
	return &Ability207{client}
}

/*
获取奇门用户列表
*/
func (ability *Ability207) TaobaoQimenTradeUsersGet(req *request.TaobaoQimenTradeUsersGetRequest) (*response.TaobaoQimenTradeUsersGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability207 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.qimen.trade.users.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoQimenTradeUsersGetResponse{}
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
删除奇门订单链路用户
*/
func (ability *Ability207) TaobaoQimenTradeUserDelete(req *request.TaobaoQimenTradeUserDeleteRequest, session string) (*response.TaobaoQimenTradeUserDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability207 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.qimen.trade.user.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoQimenTradeUserDeleteResponse{}
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
添加奇门订单链路用户
*/
func (ability *Ability207) TaobaoQimenTradeUserAdd(req *request.TaobaoQimenTradeUserAddRequest, session string) (*response.TaobaoQimenTradeUserAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability207 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.qimen.trade.user.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoQimenTradeUserAddResponse{}
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
批量发送奇门事件
*/
func (ability *Ability207) TaobaoQimenEventsProduce(req *request.TaobaoQimenEventsProduceRequest, session string) (*response.TaobaoQimenEventsProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability207 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.qimen.events.produce", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoQimenEventsProduceResponse{}
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
发出奇门事件
*/
func (ability *Ability207) TaobaoQimenEventProduce(req *request.TaobaoQimenEventProduceRequest, session string) (*response.TaobaoQimenEventProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability207 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.qimen.event.produce", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoQimenEventProduceResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
