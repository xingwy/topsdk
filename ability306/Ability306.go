package ability306

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability306/request"
	"github.com/xingwy/topsdk/ability306/response"
	"github.com/xingwy/topsdk/util"
)

type Ability306 struct {
	Client *topsdk.TopClient
}

func NewAbility306(client *topsdk.TopClient) *Ability306 {
	return &Ability306{client}
}

/*
根据子账号登录名后缀模糊搜索子账号列表
*/
func (ability *Ability306) TaobaoSubusersSubaccountSearch(req *request.TaobaoSubusersSubaccountSearchRequest, session string) (*response.TaobaoSubusersSubaccountSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability306 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subusers.subaccount.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubusersSubaccountSearchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
