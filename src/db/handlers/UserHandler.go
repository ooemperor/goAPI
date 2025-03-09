package handlers

import (
	"goAPI/src/db/models"
)

type UserHandler struct {
	*BaseHandler
}

var baseHandler = NewBaseHandler([]models.User{})
var UserHandlerObject = UserHandler{&baseHandler}
