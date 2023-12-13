package configure

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func New() *conf { return &conf{viper: viper.New(), onChangeNotifies: make(map[string][]NotifyFunc)} }

func (c *conf) AddConfigPath(in ...string) *conf {
	for _, v := range in {
		c.viper.AddConfigPath(v)
	}

	return c
}

func (c *conf) SetConfigName(in string) *conf {
	c.viper.SetConfigName(in)
	return c
}

func (c *conf) SetConfigFile(f string) *conf {
	c.viper.SetConfigFile(f)
	return c
}

func (c *conf) SetDefault(key string, value any) *conf {
	c.viper.SetDefault(key, value)
	return c
}

func (c *conf) AutomaticEnv() *conf {
	c.viper.AutomaticEnv()
	return c
}

func (c *conf) Read() (*conf, error) {
	// 从本地配置读取
	if err := c.viper.ReadInConfig(); err != nil {
		return nil, err
	}

	c.viper.WatchConfig()
	c.viper.OnConfigChange(c.OnViperChanged)

	// 初始化Nacos客户端
	if err := c.createNacosClient(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *conf) OnViperChanged(in fsnotify.Event) {

}

func (c *conf) createNacosClient() error {
	opt := c.getNacosOptions()
	if opt == nil {
		return nil
	}

	nc, err := createNacosClient(opt.Address, opt.Port, opt.Namespace, opt.Group)
	if err != nil {
		return err
	}

	c.nacos = &nacos{client: nc, group: opt.Group, onChanged: c.onNacosChanged}

	return nil
}

func (c *conf) onNacosChanged(dataId, data string) {
	if fs, ok := c.onChangeNotifies[dataId]; ok {
		for _, f := range fs {
			if f != nil {
				f(c)
			}
		}
	}
}

func (c *conf) getNacosOptions() *NacosOptions {
	if c.nacosOptions != nil {
		return c.nacosOptions
	}

	if c.useNacos {
		return &NacosOptions{
			Address:   c.GetString(AddressKey, "127.0.0.1"),
			Port:      c.GetInt(PortKey, 8848),
			Namespace: c.GetString(Namespace, ""),
			Group:     c.GetString(Group, ""),
		}
	}

	return nil
}
