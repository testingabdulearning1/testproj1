package main

import (
	"school-management-app/common/config"
	"school-management-app/common/utils"
	admin_cmd "school-management-app/src/admins/cmd"
	user_cmd "school-management-app/src/core/cmd"
	quiz_cmd "school-management-app/src/quiz/cmd"
)

func main() {
	// Load config
	utils.ImportEnv()
	config.LoadConfig()

	if config.SERVER_TYPE == "user" {
		user_cmd.RunUsersServer()
	} else if config.SERVER_TYPE == "admin" {
		admin_cmd.RunAdminsServer()
	} else if config.SERVER_TYPE == "quiz" {
		quiz_cmd.RunQuizServer()
	} else {
		user_cmd.RunUsersServer()
	}
}
