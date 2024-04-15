package request


type TaobaoJdsHluserUpdateRequest struct {
    /*
        回流信息是否开通买家端展示,可选值open,close     */
    OpenForBuyer  *string `json:"open_for_buyer,omitempty" required:"false" `
    /*
        为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS     */
    OpenNodes  *string `json:"open_nodes,omitempty" required:"false" `
}

func (s *TaobaoJdsHluserUpdateRequest) SetOpenForBuyer(v string) *TaobaoJdsHluserUpdateRequest {
    s.OpenForBuyer = &v
    return s
}
func (s *TaobaoJdsHluserUpdateRequest) SetOpenNodes(v string) *TaobaoJdsHluserUpdateRequest {
    s.OpenNodes = &v
    return s
}

func (req *TaobaoJdsHluserUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OpenForBuyer != nil) {
        paramMap["open_for_buyer"] = *req.OpenForBuyer
    }
    if(req.OpenNodes != nil) {
        paramMap["open_nodes"] = *req.OpenNodes
    }
    return paramMap
}

func (req *TaobaoJdsHluserUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}