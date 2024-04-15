package request

import (
        "topsdk/util"
    )

type TaobaoItemcatsAuthorizeGetRequest struct {
    /*
        需要返回的字段。目前支持有：brand.vid, brand.name, item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,xinpin_item_cat.cid, xinpin_item_cat.name, xinpin_item_cat.status,xinpin_item_cat.sort_order,xinpin_item_cat.parent_cid,xinpin_item_cat.is_parent     */
    Fields  *[]string `json:"fields" required:"true" `
}

func (s *TaobaoItemcatsAuthorizeGetRequest) SetFields(v []string) *TaobaoItemcatsAuthorizeGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoItemcatsAuthorizeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoItemcatsAuthorizeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}