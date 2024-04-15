package ability641

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability641/request"
    "topsdk/ability641/response"
    "topsdk/util"
)

type Ability641 struct {
    Client *topsdk.TopClient
}

func NewAbility641(client *topsdk.TopClient) *Ability641{
    return &Ability641{client}
}

/*
    查询某个卖家的店铺优惠券领取活动
*/
func (ability *Ability641) TaobaoPromotionActivityGet(req *request.TaobaoPromotionActivityGetRequest,session string) (*response.TaobaoPromotionActivityGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.activity.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionActivityGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionActivityGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家优惠券
*/
func (ability *Ability641) TaobaoPromotionCouponsGet(req *request.TaobaoPromotionCouponsGetRequest,session string) (*response.TaobaoPromotionCouponsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.coupons.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionCouponsGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionCouponsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    搭配套餐查询
*/
func (ability *Ability641) TaobaoPromotionMealGet(req *request.TaobaoPromotionMealGetRequest,session string) (*response.TaobaoPromotionMealGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.meal.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionMealGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionMealGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
