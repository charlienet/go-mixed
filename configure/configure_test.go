package configure_test

import (
	"testing"

	"github.com/charlienet/go-mixed/configure"
	"github.com/charlienet/go-mixed/json"
	"github.com/stretchr/testify/assert"
)

func TestLoadSpecifiedFile(t *testing.T) {
	conf, err := configure.New().SetConfigFile("config.toml").Read()
	t.Log(err)

	assert.Equal(t, "192.168.2.121", conf.GetString("nacos.address", ""))
	_ = conf
}

func TestNewConfigure(t *testing.T) {

}

func TestNacos(t *testing.T) {
	conf, err := configure.
		New().
		AddConfigPath(".").
		WithNacos().
		Read()

	assert.Nil(t, err)

	t.Log(conf.GetString("nacos.address", ""))

	type redis struct {
		Addrs string
	}

	r := &redis{}

	t.Log(conf.Load("redis", r))
	t.Log(json.StructToJsonIndent(r))
}
