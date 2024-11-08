package auth

import (
	"tauth/entities"
	"tauth/models"
	authsvc "tauth/services/auth"
)

/* -------------------------------------------------------------------------- */
/*                                  REGISTER                                  */
/* -------------------------------------------------------------------------- */
func registerSuccessRes(baseRes models.BaseResponse, user entities.Users) models.BaseResponse {
	var res models.BaseResponse
	var data authsvc.UserObject

	// map the data
	data.PID = user.PID.String
	data.Email = user.PrimaryEmail
	data.CreatedAt = int(user.CreatedAt.Unix())

	res.Data = data
	res.Success = baseRes.Success
	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message

	return res
}

/* -------------------------------------------------------------------------- */
/*                               USERNAME CHECK                               */
/* -------------------------------------------------------------------------- */
func emailCheckResponse(baseRes models.BaseResponse, exist bool) models.BaseResponse {
	var res models.BaseResponse
	var data authsvc.EMailCheckResponse
	data.Exist = exist

	res.Data = data
	res.Success = baseRes.Success
	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message

	return res
}
