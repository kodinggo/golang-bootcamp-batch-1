package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetInt("env"))
	fmt.Println(viper.GetString("postgres.host"))
	fmt.Println(viper.GetString("postgres.port"))
	fmt.Println(viper.GetString("postgres.user"))
}
