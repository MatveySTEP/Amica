package main

import (
	"amica/api"
	"amica/db"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

var configDefaults = map[string]interface{}{
	"gomaxprocs": 0,
	"api_addr":   "127.0.0.1:8100",
	"redis_addr": "127.0.0.1:6379",
	"db_addr":    "postgres://127.0.0.1:5432/test",
}

func init() {
	rootCmd.Flags().String("jwt_secret", "", "jwt secret")
	rootCmd.Flags().String("api_addr", "127.0.0.1:8100", "api address")
	rootCmd.Flags().String("redis_addr", "127.0.0.1:6379", "redis address")
	rootCmd.Flags().String("db_addr", "postgres://127.0.0.1:5432/test", "database url")

	viper.BindPFlag("jwt_secret", rootCmd.Flags().Lookup("jwt_secret"))
	viper.BindPFlag("api_addr", rootCmd.Flags().Lookup("api_addr"))
	viper.BindPFlag("redis_addr", rootCmd.Flags().Lookup("redis_addr"))
	viper.BindPFlag("db_addr", rootCmd.Flags().Lookup("db_addr"))
}

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		printWelcome()

		for k, v := range configDefaults {
			viper.SetDefault(k, v)
		}

		bindEnvs := []string{
			"gomaxprocs", "api_addr", "redis_addr", "db_addr", "jwt_secret",
		}
		for _, env := range bindEnvs {
			err := viper.BindEnv(env)
			if err != nil {
				log.Fatal().Err(err).Msg("error binding env variable")
			}
		}

		if os.Getenv("GOMAXPROCS") == "" {
			if viper.IsSet("gomaxprocs") && viper.GetInt("gomaxprocs") > 0 {
				runtime.GOMAXPROCS(viper.GetInt("gomaxprocs"))
			} else {
				runtime.GOMAXPROCS(runtime.NumCPU())
			}
		}
		db.Connect()
		v := viper.GetViper()
		a := api.NewApi(api.Config{Addr: v.GetString("api_addr")})
		a.Run()
	},
}

func printWelcome() {
	fmt.Println("     _              _           \n    / \\   _ __ ___ (_) ___ __ _ \n   / _ \\ | '_ ` _ \\| |/ __/ _` |\n  / ___ \\| | | | | | | (_| (_| |\n /_/   \\_\\_| |_| |_|_|\\___\\__,_|\n                                ")
}
