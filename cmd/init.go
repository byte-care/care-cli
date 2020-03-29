package cmd

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func verifyKey(accessKey string, secretKey string) (err error) {
	// result := true
	// err := nil
	return
}

func Init(configFilePath string) *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "init the cli",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "access-key", Required: true},
			&cli.StringFlag{Name: "secret-key", Required: true},
		},
		Action: func(c *cli.Context) error {
			accessKey := c.String("access-key")
			secretKey := c.String("secret-key")

			if err := verifyKey(accessKey, secretKey); err != nil {
				return err
			}

			viper.Set("access-key", accessKey)
			viper.Set("secret-key", secretKey)

			if err := viper.WriteConfigAs(configFilePath); err != nil {
				panic(err)
			}

			fmt.Println("AK:", accessKey)
			fmt.Println("SK:", secretKey)
			return nil
		},
	}
}
