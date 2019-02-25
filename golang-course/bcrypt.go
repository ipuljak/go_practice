package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func mainBcrypt() {
	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Password - ", s)
	fmt.Println("Encrypted password - ", bs)

	password := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))

	if err != nil {
		fmt.Println("Wrong password!", err)
	}
}
