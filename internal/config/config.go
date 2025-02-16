package config

import "os"

var SecretKey = os.Getenv("SECRET_KEY")
