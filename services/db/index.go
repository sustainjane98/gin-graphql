package db

type Services struct {
}

func (s Services) User() userService {
	return userService{}
}
