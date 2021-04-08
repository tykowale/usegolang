package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	userService, err := models.NewUserService(psqlInfo)

	if err != nil {
		panic(err)
	}
	defer userService.Close()
	userService.DestructiveReset()

	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}

	if err := userService.Create(&user); err != nil {
		panic(err)
	}
	user.Name = "Updated Name"
	if err := userService.Update(&user); err != nil {
		panic(err)
	}

	foundUser, err := userService.ByEmail("michael@dundermifflin.com")

	if err != nil {
		panic(err)
	}

	fmt.Println(foundUser)

	if err := userService.Delete(foundUser.ID); err != nil {
		panic(err)
	}
	_, err = userService.ByID(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user was not deleted!")
	}
}
