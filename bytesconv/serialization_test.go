package bytesconv

import (
	"encoding/json"
	"testing"
)

type SimpleUser struct {
	FirstName string
	LastName  string
}

func TestGob(t *testing.T) {
	u := SimpleUser{FirstName: "Radomir", LastName: "Sohlich"}
	buf, err := Encode(u)
	t.Log("Gob", BytesResult(buf).Hex(), err)

	var u2 SimpleUser
	if err := Decode(buf, &u2); err != nil {
		t.Fatal(err)
	}

	jBytes, _ := json.Marshal(u2)
	t.Log("Json:", BytesResult(jBytes).Hex())

	t.Logf("%+v", u2)
}
