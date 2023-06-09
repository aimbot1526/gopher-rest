package utils

func CreateUserValidator(username string, password string) (bool, string) {

	if len(username) == 0 {
		return false, "Invalid Username"
	}

	if len(password) < 6 {
		return false, "Invalid Password. Please use more than 6 characters"
	}
	return true, "User validated successfully."
}
