package errm

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

type Any interface{}

type Monad func(error) (Any, error)

func Return(v Any) Monad {
	return func(s error) (Any, error) {
		return v, s
	}
}

func Bind(m Monad, f func(Any) Monad) Monad {
	return func(s error) (Any, error) {
		newV, newS := m(s)
		if newS != nil {
			return nil, newS
		}
		return f(newV)(newS)
	}
}

// CryptoRandRead reads v as []byte and returns Monad: error -> int, error
func CryptoRandRead(v Any) Monad {
	vBytes := v.([]byte)
	return func(s error) (Any, error) {
		return rand.Read(vBytes)
	}
}

// Base64DecodeString reads v as string and returns Monad: error -> []byte, error
func Base64DecodeString(v Any) Monad {
	vString := v.(string)
	return func(s error) (Any, error) {
		return base64.StdEncoding.DecodeString(vString)
	}
}

// JSONUnmarshal reads v as []byte and returns Monad: error -> map[string]interface{}, error
func JSONUnmarshal(v Any) Monad {
	vBytes := v.([]byte)
	return func(s error) (Any, error) {
		resultMap := make(map[string]interface{})
		err := json.Unmarshal(vBytes, &resultMap)
		return resultMap, err
	}
}

// ReadFile reads v as string and returns Monad: error -> []byte, error
func ReadFile(filename Any) Monad {
	filenameString := filename.(string)
	return func(error) (Any, error) {
		return ioutil.ReadFile(filenameString)
	}
}

// BytesToStr reads v as []byte and returns Monad: error -> string, error
func BytesToStr(v Any) Monad {
	vBytes := v.([]byte)
	return func(error) (Any, error) {
		return string(vBytes), nil
	}
}
