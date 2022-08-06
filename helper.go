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


func version() string {
	return fmt.Sprintf("%s v%s", APP_NAME, APP_VERSION)
} // version


func addr() string {
	return fmt.Sprintf("%s:%s", *host, *port)
} // addr


func userKey(email string) string {
	return hash(email+SALT)
} // userKey


func hash(c string) string {

	digest := hmac.New(sha256.New, []byte(SALT))

	digest.Write([]byte(c))

	res := hex.EncodeToString(digest.Sum(nil))

	return res[:15]

} // hash


func encrypt(clearText string) string {
  return STRING_EMPTY
} // encrypt


func decrypt(cipher string) string {
  return STRING_EMPTY
} // decrypt


func initRating(max int) int {
	
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
  
	return r.Intn(max)

} // initRating


func initRangeRating(min int, max int) int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	if min > max {
		return r.Intn(min-max) + max
	} else {
		return r.Intn(max-min) + min
	}

} // initRangeRating


func initFamily() string {

	families, err := rds.SMembers(ctx, KEY_FAMILIES).Result()

	if err != nil {
		
		log.Println(err)
		return STRING_EMPTY

	} else {
		
		count := len(families)
		return families[initRating(count)]
	
	}

} // initFamily
