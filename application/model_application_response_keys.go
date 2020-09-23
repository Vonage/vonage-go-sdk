/*
Handmade file to fix the AllOf that misses the keys
*/

package application

// ApplicationResponseKeys
type ApplicationResponseKeys struct {
	PublicKey  string `json:"public_key,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
}
