package request


type TmallProductSchemaMatchRequest struct {
    /*
        商品发布的目标类目，必须是叶子类目     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。     */
    Propvalues  *string `json:"propvalues" required:"true" `
}

func (s *TmallProductSchemaMatchRequest) SetCategoryId(v int64) *TmallProductSchemaMatchRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallProductSchemaMatchRequest) SetPropvalues(v string) *TmallProductSchemaMatchRequest {
    s.Propvalues = &v
    return s
}

func (req *TmallProductSchemaMatchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.Propvalues != nil) {
        paramMap["propvalues"] = *req.Propvalues
    }
    return paramMap
}

func (req *TmallProductSchemaMatchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}