package converter

import (
	"example/graph/model"
	"example/models"
	"strconv"
)

type userConverter struct {
}

func User() userConverter {
	return userConverter{}
}

func (_ userConverter) CreateDtoToDao(c model.CreateUserDto) (*models.UserDao, error) {

	t := c.Birthdate.Time

	dao := models.UserDao{Age: c.Age, Username: c.Username, Password: c.Password, Birthdate: t, Email: c.Email, Firstname: c.Firstname, Lastname: c.Lastname}

	return &dao, nil
}

func (_ userConverter) DaoToDto(c models.UserDao) model.UserDto {
	return model.UserDto{Age: c.Age, Username: c.Username, ID: strconv.Itoa(int(c.ID)), Firstname: c.Firstname, Lastname: c.Lastname, Email: c.Email, Birthdate: model.Date{Time: c.Birthdate}}
}
