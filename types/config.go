package types

import (
	"github.com/BurntSushi/toml"
)

type configStruct struct {
	Title string
	Grpc  grpcStruct
	Eth   ethStruct
	Mysql mysqlStruct
}

type grpcStruct struct {
	GrpcAddress      string `toml:"grpc_address"`
	ValidatorAddress string `toml:"validator_address"`
}

type ethStruct struct {
	Etherscan_Mainnet_Endpoint string `toml:"etherscan_mainnet_endpoint"`
	Etherscan_Key              string `toml:"etherscan_key"`
}

type mysqlStruct struct {
	DbHost string `toml:"dbhost"`
	DbUser string `toml:"dbuser"`
	DbPass string `toml:"dbpass"`
	DbName string `toml:"dbname"`
}

func GetConfig() configStruct {
	var config configStruct
	_, err := toml.DecodeFile("config/config.toml", &config)
	if err != nil {
		panic(err)
	}

	return config
}
