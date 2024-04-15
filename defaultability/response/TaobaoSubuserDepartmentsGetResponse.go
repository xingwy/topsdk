package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoSubuserDepartmentsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   部门信息
	*/
	Departments []domain.TaobaoSubuserDepartmentsGetDepartment `json:"departments,omitempty" `
}
