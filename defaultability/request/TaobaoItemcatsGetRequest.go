package request

import (
        "topsdk/util"
    )

type TaobaoItemcatsGetRequest struct {
    /*
        商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)     */
    Cids  *[]int64 `json:"cids,omitempty" required:"false" `
    /*
        无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化     */
    Datetime  *util.LocalTime `json:"datetime,omitempty" required:"false" `
    /*
        需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 defalutValue��cid,parent_cid,name,is_parent    */
    Fields  *[]string `json:"fields,omitempty" required:"false" `
    /*
        父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)     */
    ParentCid  *int64 `json:"parent_cid,omitempty" required:"false" `
}

func (s *TaobaoItemcatsGetRequest) SetCids(v []int64) *TaobaoItemcatsGetRequest {
    s.Cids = &v
    return s
}
func (s *TaobaoItemcatsGetRequest) SetDatetime(v util.LocalTime) *TaobaoItemcatsGetRequest {
    s.Datetime = &v
    return s
}
func (s *TaobaoItemcatsGetRequest) SetFields(v []string) *TaobaoItemcatsGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItemcatsGetRequest) SetParentCid(v int64) *TaobaoItemcatsGetRequest {
    s.ParentCid = &v
    return s
}

func (req *TaobaoItemcatsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Cids != nil) {
        paramMap["cids"] = util.ConvertBasicList(*req.Cids)
    }
    if(req.Datetime != nil) {
        paramMap["datetime"] = *req.Datetime
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.ParentCid != nil) {
        paramMap["parent_cid"] = *req.ParentCid
    }
    return paramMap
}

func (req *TaobaoItemcatsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}