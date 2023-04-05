package Entity

import (
	"awesomeProject/Login/Dao"
	"fmt"
)

func ValidateUser(username string, password string, userRepository Dao.UserRepository) (bool, error) {
	passwordResult, err := userRepository.GetUser(username)

	if err != nil {
		return false, fmt.Errorf("error checking user:%s", err)

	}
	if passwordResult == password {
		return true, nil
	}
	return false, nil

}
