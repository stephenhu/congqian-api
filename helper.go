package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


func version() string {
	return fmt.Sprintf("%s v%s", APP_NAME, APP_VERSION)
} // version


func addr() string {
	return fmt.Sprintf("%s:%s", *host, *port)
} // addr


func userKey(email string) string {
	return fmt.Sprintf("%s:%s", APP_NAME, hash(email))
} // userKey


func hash(c string) string {

	digest := hmac.New(sha256.New, []byte(DEFAULT_SALT))

	digest.Write([]byte(c))

	res := hex.EncodeToString(digest.Sum(nil))

	return res[:15]

} // hash
