package dto

import "github.com/sunyd/go-demo1/model"

type UserDto struct {
	Username  string `json:"username"`
	Telephone string `json:"telphone"`
}

/**
将user对象转换为userDto对象
*/
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username:  user.Username,
		Telephone: user.Telephone,
	}
}
