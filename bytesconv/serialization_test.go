package bytesconv

import (
	"encoding/hex"
	"encoding/json"
	"testing"
	"time"
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

type delayTask struct {
	message string
	delay   time.Time
	execute func()
}

func TestMarshal(t *testing.T) {
	d := delayTask{
		message: "sssssssss",
	}

	b, err := Encode(d)
	t.Log(hex.EncodeToString(b), err)
}
