package main

import (
	"fmt"
	"log"
)


type CharAttributes struct {
	IsFemale        		bool      	`json:"isFemale"`
	Age           			int         `json:"age"`
	Height        			int         `json:"height"`
	Family        			string      `json:"family"`
	Kingdom         		string      `json:"kingdom"`
	Experience      		int         `json:"experience"`
	Luck            		int         `json:"luck"`
	Virtue      				int         `json:"virtue"`
	Loyalty         		int         `json:"loyalty"`
	Popularity    			int         `json:"popularity"`
	Health          		int         `json:"health"`
	Energy       				int         `json:"energy"`
	Intelligence    		int         `json:"intelligence"`
	Wisdom       				int         `json:"wisdom"`
	Integrity       		int         `json:"integrity"`
	Dexterity       		int         `json:"dexterity"`
	Charisma        		int         `json:"charisma"`
	Attractiveness    	int         `json:"attractiveness"`
	Strength          	int         `json:"strength"`
	Stamina           	int         `json:"stamina"`
	Balance            	int         `json:"balance"`
	Temperance          int         `json:"temperance"`
	Recovery            int         `json:"recovery"`
	Innovation          int         `json:"innovation"`
}


type MilitaryAttribute struct {
	Size								int 				`json:"size"`
	Rating          		int         `json:"rating"`
}


type MilitaryAttributes struct {
	Dagger				MilitaryAttribute		`json:"dagger"`
	Sword					MilitaryAttribute		`json:"sword"`
	Scimitar			MilitaryAttribute		`json:"scimitar"`
	Spear					MilitaryAttribute		`json:"spear"`
	HookedSpear		MilitaryAttribute		`json:"hookedSpear"`
	Staff					MilitaryAttribute		`json:"staff"`
	DragonBlade		MilitaryAttribute		`json:"dragonBlade"`
	Club					MilitaryAttribute		`json:"club"`
	Hammer				MilitaryAttribute		`json:"hammer"`
	Axe						MilitaryAttribute		`json:"axe"`
	Halberd				MilitaryAttribute		`json:"halberd"`
	Bow						MilitaryAttribute		`json:"bow"`
	CrossBow			MilitaryAttribute		`json:"crossbow"`
	HorseBow			MilitaryAttribute		`json:"horseBow"`
	Horse					MilitaryAttribute		`json:"horse"`
	Chariot				MilitaryAttribute		`json:"chariot"`
	Shield				MilitaryAttribute		`json:"shield"`
}


type User struct {
	ID 						string			`json:"id"`
	Email         string      `json:"email"`
	Pass 					string			`json:"password"`
	Name          string      `json:"name"`
	Attributes    CharAttributes
	Combat      	MilitaryAttributes
}


const (
	FIELD_AGE             = "age"
	FIELD_EMAIL						= "email"
	FIELD_FAMILY          = "family"
	FIELD_FEMALE          = "female"
	FIELD_HEIGHT          = "height"
	FIELD_NAME						= "name"
	FIELD_PASSWORD				= "password"
)


// TODO: check if email exists
func registerUser(email string, password string, name string) {

	key := fmt.Sprintf("%s:%s", KEY_USER, name)

	val, err := rds.HGet(ctx, key, FIELD_EMAIL).Result()

	if err != nil {

		log.Println(err)

		err = rds.HSet(ctx, key, FIELD_NAME, name,
			FIELD_PASSWORD, hash(password)).Err()

		if err != nil {
			log.Println(err)
		} 

	} else {
		log.Println(val)
	}

} // registerUser


func createCharacter(name string, female bool) {

	key := fmt.Sprintf("%s:%s", KEY_USER, name)

	val, err := rds.HSet(ctx, key,
		FIELD_AGE, initRangeRating(DEFAULT_AGE_MIN, DEFAULT_AGE_MAX),
    FIELD_FEMALE, female,
	  FIELD_HEIGHT, initRangeRating(DEFAULT_HEIGHT_MIN, DEFAULT_HEIGHT_MAX),
		FIELD_FAMILY, initFamily()).Result()

	if err != nil {
		log.Println(err)
	} else {
		log.Println(val)
	}

} // createCharacter
