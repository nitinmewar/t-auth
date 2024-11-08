package auth

import authsvc "tauth/services/auth"

type authHandler struct {
	authSvc authsvc.Interface
}

func Handler(authSvc authsvc.Interface) *authHandler {
	return &authHandler{
		authSvc: authSvc,
	}
}
