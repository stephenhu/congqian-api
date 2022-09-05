package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"
)


var rng *rand.Rand


func version() string {
	return fmt.Sprintf("%s v%s", APP_NAME, APP_VERSION)
} // version


func addr() string {
	return fmt.Sprintf("%s:%s", *host, *port)
} // addr


func loadConfigs(key string) {

	var err error

	if key == KEY_CONF_GLOBAL {
		err = rds.HGetAll(ctx, key).Scan(&confGlobal)
	} else if key == KEY_CONF_INIT {
		err = rds.HGetAll(ctx, key).Scan(&confInit)
	} else {
		log.Fatal("Unsupported key, aborting app... ", key)
	}

	if err != nil {
		log.Println(err)
	}
	
} // loadConfigs


func userKey(email string) string {
	return hash(email+SALT, DEFAULT_HASH_LENGTH)
} // userKey


func hash(c string, length int) string {

	if length < 1 {
		
		log.Println(ERR_HASH_INVALID_LENGTH)
		length = DEFAULT_HASH_LENGTH

	}

	digest := hmac.New(sha256.New, []byte(SALT))

	digest.Write([]byte(c))

	res := hex.EncodeToString(digest.Sum(nil))

	return res[:length]

} // hash


// TODO
func encrypt(clearText string) string {
  return STRING_EMPTY
} // encrypt


func decrypt(cipher string) string {
  return STRING_EMPTY
} // decrypt


func initRng() {

	if rng == nil {
		
		s := rand.NewSource(time.Now().UnixNano())
		rng = rand.New(s)

	}

} // initRng


func initRating(max int) int {
	return rng.Intn(max)
} // initRating


func initRangeRating(min int, max int) int {

	if min > max {
		return rng.Intn(min-max) + max
	} else {
		return rng.Intn(max-min) + min
	}

} // initRangeRating


func initFromSet(key string) string {

	s, err := rds.SMembers(ctx, key).Result()

	if err != nil {

		log.Println(err)
		return STRING_EMPTY

	} else {
		
		count := len(s)

		if count == 0 {

			log.Println(ERR_SET_NOT_FOUND)
			return STRING_EMPTY

		}	else {
			return s[initRating(count)]
		}

	}

} // initFromSet


func initMunicipal(kingdom string) string {

	municipals, err := rds.SMembers(ctx, fmt.Sprintf("%s:%s:%s", KEY_KINGDOM,
	  kingdom, KEY_MUNICIPALS)).Result()

	if err != nil {
		
		log.Println(err)
		return STRING_EMPTY

	} else {

		count := len(municipals)

		if count == 0 {

			log.Println(ERR_NO_MUNICIPALS_FOUND)
			return STRING_EMPTY

		}	else {
			return municipals[initRating(count)]
		}
	
	}

} // initMunicipal


func initGender() int {

	initRng()

	return initRating(DEFAULT_GENDER)

} // initGender
