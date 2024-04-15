package ability306

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability306/request"
    "topsdk/ability306/response"
    "topsdk/util"
)

type Ability306 struct {
    Client *topsdk.TopClient
}

func NewAbility306(client *topsdk.TopClient) *Ability306{
    return &Ability306{client}
}

/*
    根据子账号登录名后缀模糊搜索子账号列表
*/
func (ability *Ability306) TaobaoSubusersSubaccountSearch(req *request.TaobaoSubusersSubaccountSearchRequest,session string) (*response.TaobaoSubusersSubaccountSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability306 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.subaccount.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersSubaccountSearchResponse{}
    if(err != nil){
        log.Println("taobaoSubusersSubaccountSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
