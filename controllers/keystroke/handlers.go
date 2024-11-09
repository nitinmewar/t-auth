package keystroke

import "tauth/services/keystrokesvc"

type keyStrokeHandler struct {
	keystrokeSvc keystrokesvc.Interface
}

func Handler(keystrokeSvc keystrokesvc.Interface) *keyStrokeHandler {
	return &keyStrokeHandler{
		keystrokeSvc: keystrokeSvc,
	}
}
