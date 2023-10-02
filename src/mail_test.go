package main

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail("comunesondage@outlook.com", "mikaelsourati@gmail.com", "Test error", "HEY GUYS", "J2t@#g+Mw9zK")

	if err != nil {
		fmt.Println("Error SendMail")
	} else {
		fmt.Println("Success Sendmail")
	}
}
