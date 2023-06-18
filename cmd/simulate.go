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
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		var opts []repository.ListOption
		if localSearchDir := viper.GetString("searcher-dir"); localSearchDir != "" {
			opts = append(opts, repository.AmiiboListWithSearchDir(localSearchDir))
		}

		amiibo := repository.NewAmiiboSearcher(viper.GetString("searcher"))
		amiibos, err := amiibo.List(ctx, opts...)
		if err != nil {
			return err
		}

		pm3, err := repository.NewPM3Device(viper.GetString("key-retail"))
		if err != nil {
			return err
		}

		return usecase.NewSim(pm3).Action(ctx, amiibos)
	},
}

func init() {
	rootCmd.AddCommand(simCmd)

	simCmd.Flags().String("key-retail", "key_retail.bin", "specify the path of the key_retail.bin")
	simCmd.Flags().String("searcher", "remote", "specify the mode of the amiibo searcher")
	simCmd.Flags().String("searcher-dir", "~/.local/share/amiicli/local", "specify the dir path of the amiibo searcher")
	_ = viper.BindPFlags(simCmd.Flags())
}
