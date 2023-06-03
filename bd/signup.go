package bd

import (
	"fmt"

	"github.com/esaudevs/turtlesUser/tools"

	"github.com/esaudevs/turtlesUser/models"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Starting sign up")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.MySQLDate() + "')"
	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Successfully")
	return nil
}
