package compiledbuffer

import (
	"testing"

	"github.com/dlclark/regexp2"
)

func TestCom(t *testing.T) {
	regex, err := regexp2.Compile(`^\d{11}[;；](?!(37|38))\d{2}\d{6}$`, regexp2.None)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(regex.MatchString("14610522152;37764800"))
	t.Log(regex.MatchString("14610522152;33764800"))
}
