package configure

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

const (
	AddressKey = "Nacos.Address"
	PortKey    = "Nacos.Port"
	Namespace  = "Nacos.Namespace"
	Group      = "Nacos.Group"
)

type nacos struct {
	client    config_client.IConfigClient
	onChanged func(string, string)
	group     string
}

type NacosOptions struct {
	Address   string
	Port      int
	Namespace string
	Group     string
}

func (c *conf) WithNacosOptions(options *NacosOptions) *conf {
	c.nacosOptions = options
	return c
}

func (c *conf) WithNacos() *conf {
	c.useNacos = true
	return c
}

func (n *nacos) Load(dataId string, v any) error {
	voParam := vo.ConfigParam{
		DataId:   dataId,
		Group:    n.group,
		OnChange: n.onChange,
	}

	content, err := n.client.GetConfig(voParam)
	if err != nil {
		return err
	}

	if len(content) == 0 {
		return fmt.Errorf("parameters not configured:%s", dataId)
	}

	if err := json.Unmarshal([]byte(content), v); err != nil {
		return err
	}

	n.client.ListenConfig(voParam)
	return nil
}

func (n *nacos) onChange(namespace, group, dataId, data string) {
	n.onChanged(dataId, data)
}

func createNacosClient(addr string, port int, namespace, group string) (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{{
		IpAddr: addr,
		Port:   uint64(port),
	}}

	cc := constant.ClientConfig{
		NamespaceId:         namespace,
		TimeoutMs:           5000,
		LogDir:              "logs",
		CacheDir:            "cache",
		LogLevel:            "info",
		NotLoadCacheAtStart: true,
	}

	configClient, err := clients.CreateConfigClient(map[string]any{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		return nil, err
	}

	return configClient, nil
}
