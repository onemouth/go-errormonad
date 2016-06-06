package errm

import (
	"testing"
)

func TestReadFileBase64JSON(t *testing.T) {
	m := Return("testdata/test.base64")
	m = Bind(m, ReadFile)
	m = Bind(m, BytesToStr)
	m = Bind(m, Base64DecodeString)
	m = Bind(m, JSONUnmarshal)
	jsonMap, err := m(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jsonMap)
}
