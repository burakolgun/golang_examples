package error

import (
	"errors"
	"fmt"
)

func mailCheck(mailAddress string) (error, bool) {
	alreadyDefinedAddress := map[string]bool{"a@mail.com": true, "b@mail.com,": true, "c@mail.com": true, "d@mail.com": true}

	if alreadyDefinedAddress[mailAddress] {
		return errors.New("already defined address"), false
	}

	return nil, true
}

func Example() {
	mails := map[int]string{1: "a@mail.com", 2: "b@mail.com", 3: "e@mail.com"}

	for key, value := range mails {
		err, success := mailCheck(value)
		if success {
			fmt.Println(success, "success", key)
		} else {
			fmt.Println(err, key)
		}
	}

}
