/*
Copyright © 2021 Cedric L'homme

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/rangzen/eaopt-portfolio-allocation/sim"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var nbGeneration uint
var targetValue float64
var maxShares uint

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eaopt-portfolio-allocation data-file.json",
	Short: "Portfolio allocation using Genetic Algorithm.",
	Long: `Portfolio allocation with multiple targets
using eaopt for the GA part.

For example, calculate shares and targets from data-example.json
with 10000 generations and a total value of 50000:

eaopt-porfolio-allocation data-example.json -g 10000 -v 50000`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a JSON data file")
		}
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return errors.New("data file does not exist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		s, err := sim.New(args[0], nbGeneration, targetValue, maxShares)
		if err != nil {
			fmt.Println(err)
			return
		}
		s.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eaopt-portfolio-allocation.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().UintVarP(&nbGeneration, "nb-gen", "g", 10000, "number of generation")
	rootCmd.PersistentFlags().Float64VarP(&targetValue, "target-value", "v", 50000., "target value")
	rootCmd.PersistentFlags().UintVarP(&maxShares, "max-shares", "m", 10000, "maximum of shares for one value")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".eaopt-portfolio-allocation" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".eaopt-portfolio-allocation")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
