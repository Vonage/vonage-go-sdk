package vonage

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		name           string
		method         SignMethod
		params         map[string]interface{}
		secret         string
		expectedResult string
		errorExpected  bool
	}{
		{
			name:          "invalid method",
			method:        "random method",
			errorExpected: true,
		},
		{
			name:           "empty params",
			method:         MD5HASH,
			params:         map[string]interface{}{},
			secret:         "secret",
			expectedResult: "5ebe2294ecd0e0f08eab7690d2a6ee69",
		},
		{
			name:   "only sig param",
			method: MD5HASH,
			params: map[string]interface{}{
				"sig": "signature",
			},
			secret:         "secret",
			expectedResult: "5ebe2294ecd0e0f08eab7690d2a6ee69",
		},
		{
			name:   "custom param",
			method: MD5HASH,
			params: map[string]interface{}{
				"from": "NEXMO",
			},
			secret:         "secret",
			expectedResult: "2cdd20a2a0f7270545a98b3ccb87ba51",
		},
		{
			name:           "empty params but with md5 hmac",
			method:         MD5HMAC,
			params:         map[string]interface{}{},
			secret:         "secret",
			expectedResult: "5c8db03f04cec0f43bcb060023914190",
		},
		{
			name:           "empty params but with sha1 hmac",
			method:         SHA1HMAC,
			params:         map[string]interface{}{},
			secret:         "secret",
			expectedResult: "25af6174a0fcecc4d346680a72b7ce644b9a88e8",
		},
		{
			name:           "empty params but with sha256 hmac",
			method:         SHA256HMAC,
			params:         map[string]interface{}{},
			secret:         "secret",
			expectedResult: "f9e66e179b6747ae54108f82f8ade8b3c25d76fd30afde6c395822c530196169",
		},
		{
			name:           "empty params but with sha512 hmac",
			method:         SHA512HMAC,
			params:         map[string]interface{}{},
			secret:         "secret",
			expectedResult: "b0e9650c5faf9cd8ae02276671545424104589b3656731ec193b25d01b07561c27637c2d4d68389d6cf5007a8632c26ec89ba80a01c77a6cdd389ec28db43901",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GenerateSign(tc.method, tc.secret, tc.params)
			if tc.errorExpected {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Error("expected no error but got one")
				}

				h, _ := hex.DecodeString(tc.expectedResult)

				if !bytes.Equal(result, h) {
					t.Error("invalid sign generated")
				}
			}
		})
	}
}
