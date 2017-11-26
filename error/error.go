package main

import(
	"fmt"
	"errors"	
)

func mailCheck(mailAdress string) (error, bool) {
	alreadyAdress := map[string]bool{"a@mail.com": true, "b@mail.com,": true, "c@mail.com": true, "d@mail.com": true}

	if alreadyAdress[mailAdress] {
		return errors.New("Already adress"), false
	}

	return nil, true
}

func main() {
	mails := map[int]string{1: "a@mail.com", 2: "b@mail.com", 3: "e@mail.com"}
	
	for key, value := range mails {
		error, success := mailCheck(value)
		if success {
			fmt.Println(success, "success", key)
		} else {
			fmt.Println(error, key)
		}
	}
	
}