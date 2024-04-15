package request


type TmallProductAddSchemaGetRequest struct {
    /*
        商品发布的目标类目，必须是叶子类目     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        品牌ID     */
    BrandId  *int64 `json:"brand_id,omitempty" required:"false" `
}

func (s *TmallProductAddSchemaGetRequest) SetCategoryId(v int64) *TmallProductAddSchemaGetRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallProductAddSchemaGetRequest) SetBrandId(v int64) *TmallProductAddSchemaGetRequest {
    s.BrandId = &v
    return s
}

func (req *TmallProductAddSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    return paramMap
}

func (req *TmallProductAddSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}