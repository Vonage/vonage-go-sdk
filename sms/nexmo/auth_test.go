package nexmo

import (
	"testing"
)

func TestGetCreds(*testing.T) {
}

func TestCreateAuthFromKeySecret(*testing.T) {
	// create the auth struct
	myAuth := CreateAuthFromKeySecret("123", "456")

	// check it does what we expect
	myAuth.getCreds()
}
