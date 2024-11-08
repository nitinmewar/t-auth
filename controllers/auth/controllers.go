package auth

import (
	"net/http"
	"tauth/tautherrors"
	"tauth/utils"

	"github.com/gin-gonic/gin"
)

/* ------------------------------- EMAIL CHECK ------------------------------ */
func (h *authHandler) EmailCheck(ctx *gin.Context) {
	req, err := validateEmailCheckReq(ctx)
	if err != nil {
		tautherrors.Validation(ctx, err.Error())
		return
	}

	res, exist, err := h.authSvc.EmailCheck(ctx, req)
	if err != nil {
		tautherrors.InternalServer(ctx, err.Error())
		return
	}

	result := emailCheckResponse(res, exist)
	utils.ReturnJSONStruct(ctx, result)
}

/* --------------------------------- SINGUP --------------------------------- */
func (h *authHandler) Signup(ctx *gin.Context) {
	req, err := validateRegisterReq(ctx)
	if err != nil {
		tautherrors.Validation(ctx, err.Error())
		return
	}

	res, user, err := h.authSvc.RegisterUser(ctx, req)
	if err != nil {
		tautherrors.InternalServer(ctx, err.Error())
		return
	}

	result := registerSuccessRes(res, user)
	utils.ReturnJSONStruct(ctx, result)
}

/* ---------------------------------- LOGIN --------------------------------- */
func (h *authHandler) Login(ctx *gin.Context) {
	req, err := validateLoginReq(ctx)
	if err != nil {
		tautherrors.Validation(ctx, err.Error())
		return
	}

	res, user, err := h.authSvc.LoginUser(ctx, req)
	if err != nil {
		tautherrors.InternalServer(ctx, err.Error())
		return
	}

	if res.StatusCode != http.StatusOK {
		tautherrors.HandleServiceCodes(ctx, res)
		return
	}

	result := registerSuccessRes(res, user)
	utils.ReturnJSONStruct(ctx, result)
}
