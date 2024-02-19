package initialization

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"2106A-zg6/orderAndGoods/goodsproject/order/gloabl"
	"2106A-zg6/orderAndGoods/goodsproject/order/internal/config"
)

func InitNa(c config.Config) {

	fmt.Println(c.NacInfo)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: c.NacInfo.Host,
			Port:   uint64(c.NacInfo.Port),
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         c.NacInfo.Space, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: "partice.yaml",
		Group:  "devs"})

	json.Unmarshal([]byte(content), &gloabl.Config)

	fmt.Println("nacos：", gloabl.Config)

}
