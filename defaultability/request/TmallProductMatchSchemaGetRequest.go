package request


type TmallProductMatchSchemaGetRequest struct {
    /*
        商品发布的目标类目，必须是叶子类目     */
    CategoryId  *int64 `json:"category_id" required:"true" `
}

func (s *TmallProductMatchSchemaGetRequest) SetCategoryId(v int64) *TmallProductMatchSchemaGetRequest {
    s.CategoryId = &v
    return s
}

func (req *TmallProductMatchSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    return paramMap
}

func (req *TmallProductMatchSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}