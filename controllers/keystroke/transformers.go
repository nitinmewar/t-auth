package keystroke

import (
	"tauth/entities"
	"tauth/models"
	authsvc "tauth/services/auth"
)

/* -------------------------------------------------------------------------- */
/*                                  REGISTER                                  */
/* -------------------------------------------------------------------------- */
func keystrokeSuccessRes(baseRes models.BaseResponse, user entities.Users) models.BaseResponse {
	var res models.BaseResponse
	var data authsvc.UserObject

	// map the data
	data.PID = user.PID.String
	data.Email = user.PrimaryEmail
	data.CreatedAt = int(user.CreatedAt.Unix())
	data.IsKeystrokeDone = user.IsKeystrokeCalculated

	res.Data = data
	res.Success = baseRes.Success
	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message

	return res
}
