package configure

import (
	"github.com/charlienet/go-mixed/expr"
	"github.com/spf13/viper"
)

type Configure interface {
	GetString(string, string) string
}

type NotifyFunc func(Configure)

type conf struct {
	viper            *viper.Viper            //
	nacos            *nacos                  //
	onChangeNotifies map[string][]NotifyFunc // 已经注册的配置变更通知
	nacosOptions     *NacosOptions           //
	useNacos         bool                    //
}

func (c *conf) GetString(key string, defaultValue string) string {
	if c.viper.IsSet(key) {
		return c.viper.GetString(key)
	}

	return defaultValue
}

func (c *conf) GetInt(key string, defaultValue int) int {
	return expr.Ternary(c.viper.IsSet(key), c.viper.GetInt(key), defaultValue)
}

func (c *conf) Load(dataId string, v any, onChanged ...NotifyFunc) error {
	if err := c.nacos.Load(dataId, v); err != nil {
		return err
	}

	if len(onChanged) > 0 {
		c.onChangeNotifies[dataId] = onChanged
	}

	return nil
}
