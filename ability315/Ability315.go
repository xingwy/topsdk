package ability315

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability315/request"
	"github.com/xingwy/topsdk/ability315/response"
	"github.com/xingwy/topsdk/util"
)

type Ability315 struct {
	Client *topsdk.TopClient
}

func NewAbility315(client *topsdk.TopClient) *Ability315 {
	return &Ability315{client}
}

/*
rds创建数据库账户
*/
func (ability *Ability315) TaobaoRdsDbCreateaccount(req *request.TaobaoRdsDbCreateaccountRequest) (*response.TaobaoRdsDbCreateaccountResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability315 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.rds.db.createaccount", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoRdsDbCreateaccountResponse{}
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
rds获取RDS的DB
*/
func (ability *Ability315) TaobaoRdsDbGetdb(req *request.TaobaoRdsDbGetdbRequest) (*response.TaobaoRdsDbGetdbResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability315 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.rds.db.getdb", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoRdsDbGetdbResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
