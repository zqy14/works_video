package init

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"video1/video-api/basic/config"
)

func init() {
	v := viper.New()
	v.SetConfigFile("./dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		return
	}
	NaCosConfig := config.NaCos{}
	if err := v.Unmarshal(&NaCosConfig); err != nil {
		return
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         NaCosConfig.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      NaCosConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        NaCosConfig.Port,
			Scheme:      "http",
		},
	}

	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: NaCosConfig.DataId,
		Group:  NaCosConfig.Group})
	json.Unmarshal([]byte(content), &config.GlobalNaCos)

}
