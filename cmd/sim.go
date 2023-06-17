/*
Copyright © 2023 Ambor <saltbo@foxmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/bonaysoft/amiicli/internal/repository"
	"github.com/bonaysoft/amiicli/internal/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// simCmd represents the sim command
var simCmd = &cobra.Command{
	Use:     "simulate",
	Aliases: []string{"sim"},
	Short:   "simulate amiibos",
	Run: func(cmd *cobra.Command, args []string) {
		amiibo := repository.NewAmiiboLocal()
		pm3, err := repository.NewPM3Device(viper.GetString("key-retail"))
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := usecase.NewSim(amiibo, pm3).Action(cmd.Context()); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(simCmd)

	simCmd.Flags().String("key-retail", "key_retail.bin", "specify the path of the key_retail.bin")
	_ = viper.BindPFlags(simCmd.Flags())
}
