# go-errormonad

A state monad for error handling in golang

## Example

in monad_test.go

```golang
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
```

In the above example, usually, you should write `if err...` four times.

Howerver, by taking advatage of the monad, you only need check once in the end.



