package main

import (
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"
	"github.com/spf13/viper"
	"github.com/kan-fun/kan-sdk-go"
)

var configFilePath string
var client *kan_sdk.Client

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configFilePath = path.Join(homeDir, ".kanrc.yml")
	if _, err := os.Stat(configFilePath); err == nil {

	} else if os.IsNotExist(err) {
		panic("Please use kan-config to init the env. 😀")
	} else {
		panic(err)
	}


	viper.SetConfigName(".kanrc") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME")  // call multiple times to add many search paths

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	AccessKey := viper.GetString("access-key")
	SecretKey := viper.GetString("secret-key")

	client_local, err := kan_sdk.NewClient(AccessKey, SecretKey)
	if err != nil {
		panic(err)
	}

	client = client_local
}

func main() {
	app := &cli.App{
		Name: "kan",
		Usage: "👧💻 CLI for Kan 💻👦",
		HelpName: "kan",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "record"},
		},
		Action: index,
	}
	app.UseShortOptionHandling = true
	/*
		python main.py; kan
	*/

	// app.Commands = []*cli.Command{
	// 	cmd.Init(configFilePath),
	// }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
