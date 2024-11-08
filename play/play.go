package main

import (
	"fmt"
	"regexp"
)

func main() {
	emailref := regexp.MustCompile(`^[a-z0-9.]{1,20}@[a-z0-9.]{1,20}\.\w{2,5}$`)
	fmt.Println(emailref.MatchString("abcdefghijklmno@stud.kfu.com"))
}
