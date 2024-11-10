package keystroke

import (
	"fmt"
	"tauth/tautherrors"
	"tauth/utils"

	"github.com/gin-gonic/gin"
)

/* ---------------------------- CREATE KEYSTROKE ---------------------------- */
func (h *keyStrokeHandler) CreateKeyStroke(ctx *gin.Context) {
	req, err := validateKeystrokeReq(ctx)
	if err != nil {
		tautherrors.Validation(ctx, err.Error())
		return
	}

	res, user, err := h.keystrokeSvc.CreateKeyStroke(ctx, req)
	fmt.Println(res, err, user)
	res.Message = err.Error()
	if err != nil {
		tautherrors.InternalServer(ctx, err.Error())
		return
	}

	result := keystrokeSuccessRes(res, user)

	utils.ReturnJSONStruct(ctx, result)
}
