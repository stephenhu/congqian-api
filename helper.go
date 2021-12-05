package main

import (
	"fmt"
)


func version() string {
	return fmt.Sprintf("%s v%s", APP_NAME, APP_VERSION)
} // version
