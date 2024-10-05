package service

func NewUserService(ur userRepository) *UserService {
	return &UserService{ur: ur}
}
