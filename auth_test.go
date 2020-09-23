package vonage

import "testing"

func TestAuthGetCreds(t *testing.T) {
	// create the auth struct
	myAuth := CreateAuthFromKeySecret("123", "456")
	myCreds := myAuth.GetCreds()

	if myCreds[0] != "123" {
		t.Error("Key creds are incorrect")
	}

	if myCreds[1] != "456" {
		t.Error("Secret creds are incorrect")
	}
}

func TestAuthCreateAuthFromKeySecret(t *testing.T) {
	// create the auth struct
	myAuth := CreateAuthFromKeySecret("123", "456")

	if myAuth.apiKey != "123" {
		t.Error("Auth key not set")
	}

	if myAuth.apiSecret != "456" {
		t.Error("Auth secret not set")
	}
}
