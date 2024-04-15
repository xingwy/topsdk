package request

import (
        "topsdk/util"
    )

type TaobaoItemSkuGetRequest struct {
    /*
        需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        Sku的id。可以通过taobao.item.seller.get得到     */
    SkuId  *int64 `json:"sku_id" required:"true" `
    /*
        商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。     */
    NumIid  *int64 `json:"num_iid,omitempty" required:"false" `
}

func (s *TaobaoItemSkuGetRequest) SetFields(v []string) *TaobaoItemSkuGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItemSkuGetRequest) SetSkuId(v int64) *TaobaoItemSkuGetRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoItemSkuGetRequest) SetNumIid(v int64) *TaobaoItemSkuGetRequest {
    s.NumIid = &v
    return s
}

func (req *TaobaoItemSkuGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    return paramMap
}

func (req *TaobaoItemSkuGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}