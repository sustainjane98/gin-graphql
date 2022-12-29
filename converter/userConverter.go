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

func (_ userConverter) CreateDtoToDao(c model.CreateUserDto) models.UserDao {
	return models.UserDao{Age: c.Age, Username: c.Username, Password: c.Password}
}

func (_ userConverter) DaoToDto(c models.UserDao) model.UserDto {
	return model.UserDto{Age: c.Age, Username: c.Username, ID: strconv.Itoa(int(c.ID))}
}
