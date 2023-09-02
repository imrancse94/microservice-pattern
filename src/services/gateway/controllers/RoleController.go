package controllers

import (
	"gateway/Helper"
	"gateway/constant"
	"gateway/localize"
	"gateway/models"
	"gateway/requests"
	"gateway/response"
	"net/http"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {

	roles := models.GetRoles()

	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("success", ""),
		Data:       roles,
	}
	response.SuccessRespond(res, w)
}

func AddRole(w http.ResponseWriter, r *http.Request) {
	request := requests.AddRoleRequest{}
	Helper.Request(r, &request)

	statusCode := constant.Status("FAILED")

	if true {
		statusCode = constant.Status("SUCCESS")
	}

	res := response.Response{
		StatusCode: statusCode,
		Message:    localize.Trans("success", ""),
		Data:       request,
	}
	response.SuccessRespond(res, w)
}
