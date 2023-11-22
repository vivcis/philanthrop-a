package controllers

import "server/internal/ports"

type HTTPHandler struct {
	UserService ports.UserService
}

func NewHTTPHandler(userService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		UserService: userService,
	}
}
