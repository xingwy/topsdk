package request

import (
        "topsdk/util"
    )

type TaobaoItemSkusGetRequest struct {
    /*
        需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        sku所属商品数字id，必选。num_iid个数不能超过40个     */
    NumIids  *[]string `json:"num_iids" required:"true" `
}

func (s *TaobaoItemSkusGetRequest) SetFields(v []string) *TaobaoItemSkusGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItemSkusGetRequest) SetNumIids(v []string) *TaobaoItemSkusGetRequest {
    s.NumIids = &v
    return s
}

func (req *TaobaoItemSkusGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.NumIids != nil) {
        paramMap["num_iids"] = util.ConvertBasicList(*req.NumIids)
    }
    return paramMap
}

func (req *TaobaoItemSkusGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}