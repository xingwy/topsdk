package ability314

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability314/request"
	"github.com/xingwy/topsdk/ability314/response"
	"github.com/xingwy/topsdk/util"
)

type Ability314 struct {
	Client *topsdk.TopClient
}

func NewAbility314(client *topsdk.TopClient) *Ability314 {
	return &Ability314{client}
}

/*
获取开通的订单同步服务的用户
*/
func (ability *Ability314) TaobaoJushitaJdpUsersGet(req *request.TaobaoJushitaJdpUsersGetRequest) (*response.TaobaoJushitaJdpUsersGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability314 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jushita.jdp.users.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJushitaJdpUsersGetResponse{}
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
删除数据推送用户
*/
func (ability *Ability314) TaobaoJushitaJdpUserDelete(req *request.TaobaoJushitaJdpUserDeleteRequest) (*response.TaobaoJushitaJdpUserDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability314 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jushita.jdp.user.delete", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJushitaJdpUserDeleteResponse{}
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
添加数据推送用户
*/
func (ability *Ability314) TaobaoJushitaJdpUserAdd(req *request.TaobaoJushitaJdpUserAddRequest, session string) (*response.TaobaoJushitaJdpUserAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability314 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jushita.jdp.user.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJushitaJdpUserAddResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
