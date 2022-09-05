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


// TODO: job only holds the current one, no historical info is stored
type User struct {
	ID 						string			`json:"id"`
	Email         string      `json:"email"`
	Pass 					string			`json:"password"`
	Name          string      `json:"name"`
	Job          	string      `json:"job"`
	Municipal     string      `json:"municipal"`
	Attributes    CharAttributes
	Combat      	MilitaryAttributes
}


const (
	ATTR_AGE             		= "age"
	ATTR_ATTRACTIVENESS  		= "attractiveness"
	ATTR_BALANCE         		= "balance"
	ATTR_CHARISMA        		= "charisma"
	ATTR_DEXTERITY       		= "dexterity"
	ATTR_ELOQUENCE          = "eloquence"
	ATTR_EMAIL							= "email"
	ATTR_ENERGY          		= "energy"
	ATTR_EXPERIENCE      		= "experience"
	ATTR_FAMILY          		= "family"
	ATTR_FEMALE          		= "female"
	ATTR_HEALTH          		= "health"
	ATTR_HEIGHT          		= "height"
	ATTR_INNOVATION        	= "innovation"
	ATTR_INTEGRITY       		= "integrity"
	ATTR_INTELLIGENCE    		= "intelligence"
	ATTR_KINGDOM            = "kingdom"
	ATTR_LITERACY           = "literacy"
	ATTR_LOYALTY         		= "loyalty"
	ATTR_LUCK            		= "luck"
	ATTR_MUNICIPAL          = "municipal"
	ATTR_NAME								= "name"
	ATTR_PASSWORD						= "password"
	ATTR_POPULARITY      		= "popularity"
	ATTR_RECOVERY          	= "recovery"
	ATTR_STAMINA						= "stamina"
	ATTR_STRENGTH        		= "strength"
	ATTR_TEMPERANCE        	= "temperance"
	ATTR_VIRTUE          		= "virtue"
	ATTR_WEALTH             = "wealth"
	ATTR_WISDOM          		= "wisdom"
)


const (
	ATTR_DAGGER         		= "dagger"
	ATTR_SWORD          		= "sword"
	ATTR_SCIMITAR          	= "scimitar"
	ATTR_SPEAR          		= "spear"
	ATTR_HOOKED_SPEAR       = "hookedspear"
	ATTR_STAFF       				= "staff"
	ATTR_DRAGON_BLADE       = "dragonblade"
	ATTR_CLUB       				= "club"
	ATTR_HAMMER       			= "hammer"
	ATTR_AXE       					= "axe"
	ATTR_HALBERD       			= "halberd"
	ATTR_BOW       					= "bow"
	ATTR_CROSSBOW       		= "crossbow"
	ATTR_HORSE_BOW       		= "horsebow"
	ATTR_HORSE       				= "horse"
	ATTR_CHARIOT       			= "chariot"
	ATTR_SHIELD       			= "shield"
)


const (
	ATTR_BEAR       				= "bear"
	ATTR_CRANE       				= "crane"
	ATTR_DRAGON       			= "dragon"
	ATTR_DRUNKEN       			= "drunken"
	ATTR_IRONFIST       		= "ironfist"
	ATTR_LEOPARD       			= "leopard"
	ATTR_MANTIS       			= "mantis"
	ATTR_MONKEY       			= "monkey"
	ATTR_SNAKE       				= "snake"
	ATTR_TIGER       				= "tiger"
)


const (
	ATTR_CLEANING           = "cleaning"
	ATTR_COOKING            = "cooking"
	ATTR_CRAFTS							= "crafts"
)


func userExists(name string) bool {

	val, err := rds.SIsMember(ctx, KEY_USERS, name).Result()

	if err != nil {

		// TODO: this error is wrong
		log.Println(err)
		return false

	} else {
		return val
	}

} // userExists


func addUser(name string) {

	err := rds.SAdd(ctx, KEY_USERS, name).Err()

	if err != nil {
		log.Println(err)
	}

} // addUser


// TODO: check if email exists
func registerUser(email string, password string, name string) {

	if !userExists(name) {

		addUser(name)

		key := fmt.Sprintf("%s:%s", KEY_USER, name)

		val, err := rds.HGet(ctx, key, ATTR_EMAIL).Result()
	
		if err != nil {
	
			log.Println(err)
	
			err = rds.HSet(ctx, key, ATTR_NAME, name,
				ATTR_PASSWORD, hash(password, confGlobal.HashLength)).Err()
	
			if err != nil {
				log.Println(err)
			} 
	
		} else {
			log.Println(val)
		}

	}

} // registerUser


func createCharacter(name string, female bool) {

	if userExists(name) {

		key := fmt.Sprintf("%s:%s", KEY_USER, name)

		kingdom := initFromSet(KEY_KINGDOMS)

		err := rds.HSet(ctx, key,
			ATTR_AGE, initRangeRating(DEFAULT_AGE_MIN, DEFAULT_AGE_MAX),
			ATTR_KINGDOM, kingdom,
			ATTR_FEMALE, female,
			ATTR_HEIGHT, initRangeRating(confInit.HeightMin, confInit.HeightMax),
			ATTR_FAMILY, initFromSet(KEY_FAMILIES),
			ATTR_MUNICIPAL, initMunicipal(kingdom),
			ATTR_NAME, name,
			ATTR_EXPERIENCE, 0,
			ATTR_WEALTH, initRating(confInit.WealthMax),
			ATTR_LUCK, initRating(confInit.SkillMax),
			ATTR_VIRTUE, initRating(confInit.SkillMax),
			ATTR_LOYALTY, initRating(confInit.SkillMax),
			ATTR_POPULARITY, initRating(confInit.SkillMax),
			ATTR_HEALTH, initRating(confInit.SkillMax),
			ATTR_ENERGY, initRating(confInit.SkillMax),
			ATTR_INTELLIGENCE, initRating(confInit.SkillMax),
			ATTR_WISDOM, initRating(confInit.SkillMax),
			ATTR_INTEGRITY, initRating(confInit.SkillMax),
			ATTR_DEXTERITY, initRating(confInit.SkillMax),
			ATTR_CHARISMA, initRating(confInit.SkillMax),
			ATTR_ATTRACTIVENESS, initRating(confInit.SkillMax),
			ATTR_STRENGTH, initRating(confInit.SkillMax),
			ATTR_STAMINA, initRating(confInit.SkillMax),
			ATTR_BALANCE, initRating(confInit.SkillMax),
			ATTR_TEMPERANCE, initRating(confInit.SkillMax),
			ATTR_RECOVERY, initRating(confInit.SkillMax),
			ATTR_INNOVATION, initRating(confInit.SkillMax),
			ATTR_DAGGER, initRating(confInit.SkillMax),
			ATTR_SWORD, initRating(confInit.SkillMax),
			ATTR_SCIMITAR, initRating(confInit.SkillMax),
			ATTR_SPEAR, initRating(confInit.SkillMax),
			ATTR_HOOKED_SPEAR, initRating(confInit.SkillMax),
			ATTR_STAFF, initRating(confInit.SkillMax),
			ATTR_DRAGON_BLADE, initRating(confInit.SkillMax),
			ATTR_CLUB, initRating(confInit.SkillMax),
			ATTR_HAMMER, initRating(confInit.SkillMax),
			ATTR_AXE, initRating(confInit.SkillMax),
			ATTR_HALBERD, initRating(confInit.SkillMax),
			ATTR_BOW, initRating(confInit.SkillMax),
			ATTR_CROSSBOW, initRating(confInit.SkillMax),
			ATTR_HORSE_BOW, initRating(confInit.SkillMax),
			ATTR_HORSE, initRating(confInit.SkillMax),
			ATTR_CHARIOT, initRating(confInit.SkillMax),
			ATTR_SHIELD, initRating(confInit.SkillMax),
			ATTR_BEAR, initRating(confInit.SkillMax),
			ATTR_CRANE, initRating(confInit.SkillMax),
			ATTR_DRUNKEN, initRating(confInit.SkillMax),
			ATTR_IRONFIST, initRating(confInit.SkillMax),
			ATTR_LEOPARD, initRating(confInit.SkillMax),
			ATTR_MANTIS, initRating(confInit.SkillMax),
			ATTR_MONKEY, initRating(confInit.SkillMax),
			ATTR_SNAKE, initRating(confInit.SkillMax),
			ATTR_TIGER, initRating(confInit.SkillMax),
		).Err()
	
		if err != nil {
			log.Println(err)
		}
	
	}

} // createCharacter
