package validators

type ValidatorServices struct {
}

func (ValidatorServices) User() *userService {
	return &userService{}
}
