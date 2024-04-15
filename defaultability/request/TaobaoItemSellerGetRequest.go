package request

import (
        "topsdk/util"
    )

type TaobaoItemSellerGetRequest struct {
    /*
        需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        商品数字ID     */
    NumIid  *int64 `json:"num_iid" required:"true" `
}

func (s *TaobaoItemSellerGetRequest) SetFields(v []string) *TaobaoItemSellerGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItemSellerGetRequest) SetNumIid(v int64) *TaobaoItemSellerGetRequest {
    s.NumIid = &v
    return s
}

func (req *TaobaoItemSellerGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    return paramMap
}

func (req *TaobaoItemSellerGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}