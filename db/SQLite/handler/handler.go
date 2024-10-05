package handler

func NewUserHandler(us userService) *UserHandler {
	return &UserHandler{us: us}
}
