package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length     = flag.Int("length", 12, "Password length")
	useLower   = flag.Bool("lower", true, "Use lowercase letters")
	useUpper   = flag.Bool("upper", true, "Use uppercase letters")
	useDigits  = flag.Bool("digits", false, "Use digits")
	useSpecial = flag.Bool("special", false, "Use special characters")
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*()-_=+,.?/:;{}[]`~"
)

func generatePassword() string {
	var charset string

	if *useLower {
		charset += lowercase
	}

	if *useUpper {
		charset += uppercase
	}

	if *useDigits {
		charset += digits
	}

	if *useSpecial {
		charset += special
	}

	if charset == "" {
		panic("No character sets selected")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]byte, *length)
	for i := range password {
		password[i] = charset[r.Intn(len(charset))]
	}

	return string(password)
}

func main() {
	flag.Parse()

	if *length <= 0 {
        fmt.Println("Error: Password length must be greater than 0.")
        return
    }

	if !(*useLower || *useUpper || *useDigits || *useSpecial) {
        fmt.Println("Error: At least one character set must be selected.")
		fmt.Println("Use -h for help.")
        return
    }

	password := generatePassword()
	fmt.Println(password)
}
