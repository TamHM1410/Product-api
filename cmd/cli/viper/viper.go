package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("❌ Cannot read config: %v", err)
	}

	log.Println("✅ Config loaded from:", v.ConfigFileUsed())
	log.Printf("Config content: %+v\n", v.AllSettings())
	return v
}
