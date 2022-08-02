package main

import (
	//"fmt"
	"log"
)

type User struct {
	ID 						string			`json:"id"`
	Email         string      `json:"email"`
	Pass 				  string			`json:"password"`
	Name          string      `json:"name"`
	Gender        string      `json:"gender"`
	Age           int         `json:"age"`
	Height        int         `json:"height"`
	Family        string      `json:"family"`
}


const (
	FIELD_NAME						= "name"
	FIELD_PASSWORD				= "password"
)


func createUser(email string, password string, name string) {

	saltedId := userKey(email)

	val, err := rds.HGet(ctx, saltedId, FIELD_NAME).Result()

	if err != nil {
		log.Println(err)

		err = rds.HSet(ctx, saltedId, FIELD_NAME, name,
			FIELD_PASSWORD, hash(password)).Err()

			if err != nil {
				log.Println(err)
			} 

	} else {
		log.Println(val)
	}

} // createUser
