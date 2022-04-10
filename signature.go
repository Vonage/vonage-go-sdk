package vonage

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"sort"
	"strings"
)

type SignMethod string

const (
	MD5HASH    SignMethod = "md5hash"
	MD5HMAC    SignMethod = "md5hmac"
	SHA1HMAC   SignMethod = "sha1hmac"
	SHA256HMAC SignMethod = "sha256hmac"
	SHA512HMAC SignMethod = "sha512hmac"
)

// GenerateSign generates the signature based on https://developer.vonage.com/concepts/guides/signing-messages
// Params needs to contain all body and query parameters
func GenerateSign(method SignMethod, secret string, params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = map[string]interface{}{}
	}

	delete(params, "sig")

	keys := []string{}

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	queryToSign := ""

	for _, k := range keys {
		queryToSign += "&" + k + "=" +
			strings.NewReplacer("&", "_", "=", "_").Replace(fmt.Sprintf("%v", params[k]))
	}

	var h hash.Hash

	switch method {
	case MD5HASH:
		queryToSign += secret

		result := md5.Sum([]byte(queryToSign))

		return result[:], nil
	case MD5HMAC:
		h = hmac.New(md5.New, []byte(secret))
	case SHA1HMAC:
		h = hmac.New(sha1.New, []byte(secret))
	case SHA256HMAC:
		h = hmac.New(sha256.New, []byte(secret))
	case SHA512HMAC:
		h = hmac.New(sha512.New, []byte(secret))
	default:
		return nil, fmt.Errorf("invalid method: %s", method)
	}
	_, err := h.Write([]byte(queryToSign))
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
