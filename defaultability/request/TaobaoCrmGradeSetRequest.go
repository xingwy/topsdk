package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoCrmGradeSetRequest struct {
	/*
	        只对设置的层级等级有效，必须要在amount和count参数中选择一个<br>
	amount参数的填写规范：升级到下一个级别的需要的交易额，单位为分,必须全部填写.例如10000,20000,30000，其中10000表示非会员升级到普通的所需的交易额，20000表示普通升级到高级所需的交易额，层级等级中最高等级的下一个等级默认为0。会员等级越高，所需交易额必须越高。     */
	Amount *[]int64 `json:"amount,omitempty" required:"false" `
	/*
	   会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP     */
	Grade *[]int64 `json:"grade" required:"true" `
	/*
		        只对设置的层级等级有效，必须要在amount和count参数中选择一个<br>
		count参数的填写规范：
		升级到下一个级别的需要的交易量,必须全部填写. 以逗号分隔,例如100,200,300，其中100表示非会员升级到普通会员交易量。层级等级中最高等级的下一个等级的交易量默认为0。会员等级越高，交易量必须越高。     */
	Count *[]int64 `json:"count,omitempty" required:"false" `
	/*
	        会员级别折扣率。会员等级越高，折扣必须越低。
	950即9.5折，888折即8.88折。出于安全原因，折扣现最低只能设置到700即7折。     */
	Discount *[]int64 `json:"discount" required:"true" `
	/*
	   是否设置达到某一会员等级的交易量和交易额，必填。4个级别都需要设置，如入参为true,true,true,false时，表示设置达到高级会员、VIP会员的交易量或者交易额，不设置达到至尊会员的交易量和交易额     */
	Hierarchy *[]bool `json:"hierarchy" required:"true" `
}

func (s *TaobaoCrmGradeSetRequest) SetAmount(v []int64) *TaobaoCrmGradeSetRequest {
	s.Amount = &v
	return s
}
func (s *TaobaoCrmGradeSetRequest) SetGrade(v []int64) *TaobaoCrmGradeSetRequest {
	s.Grade = &v
	return s
}
func (s *TaobaoCrmGradeSetRequest) SetCount(v []int64) *TaobaoCrmGradeSetRequest {
	s.Count = &v
	return s
}
func (s *TaobaoCrmGradeSetRequest) SetDiscount(v []int64) *TaobaoCrmGradeSetRequest {
	s.Discount = &v
	return s
}
func (s *TaobaoCrmGradeSetRequest) SetHierarchy(v []bool) *TaobaoCrmGradeSetRequest {
	s.Hierarchy = &v
	return s
}

func (req *TaobaoCrmGradeSetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Amount != nil {
		paramMap["amount"] = util.ConvertBasicList(*req.Amount)
	}
	if req.Grade != nil {
		paramMap["grade"] = util.ConvertBasicList(*req.Grade)
	}
	if req.Count != nil {
		paramMap["count"] = util.ConvertBasicList(*req.Count)
	}
	if req.Discount != nil {
		paramMap["discount"] = util.ConvertBasicList(*req.Discount)
	}
	if req.Hierarchy != nil {
		paramMap["hierarchy"] = util.ConvertBasicList(*req.Hierarchy)
	}
	return paramMap
}

func (req *TaobaoCrmGradeSetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
