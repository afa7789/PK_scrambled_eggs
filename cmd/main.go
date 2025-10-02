package main

import (
	"fmt"
	"os"
	"time"

	"github.com/afa7789/PK_scrambled_eggs/internal/scramble"
	nconst "github.com/afa7789/PK_scrambled_eggs/pkg/natural_const"
	"github.com/spf13/cobra"
)

func main() {
	var intFlag int
	var stringFlag string
	var dateFlag string
	var fileFlag string
	var debug bool

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "An application that processes flags",
		Long: `An application that processes flags with the following options:
	-i, --constant int   An integer flag (0-5)
	-s, --string string  A string flag (default "default")
	-f, --file string    A filepath to read string input from
	-d, --date string    A date flag (MM-DD-YYYY) (default "01-03-2009")
	-v, --verbose Enable verbose output
	-h, --help           Show help message`,
		Run: func(cmd *cobra.Command, args []string) {

			// integer from 0 to 5.
			if intFlag < 0 || intFlag > 5 {
				fmt.Println("Integer flag must be between 0 and 5")
				os.Exit(1)
			}

			// 0 - Number PI π
			// 1 - Euler's number (e)
			// 2 - Golden ratio (phi), ϕ
			// 3 - Natural logarithm of 2, ln(2)
			// 4 - Euler's constant γ
			// 5 - Pythagoras' constant √2
			if !(dateFlag == "00-00-0000") {
				_, err := time.Parse("01-02-2006", dateFlag)
				if err != nil {
					fmt.Println("Date must be in the format MM-DD-YYYY")
					os.Exit(1)
				}
			}

			// If fileFlag is set, read string from file
			if fileFlag != "" {
				data, err := os.ReadFile(fileFlag)
				if err != nil {
					fmt.Printf("Failed to read file '%s': %v\n", fileFlag, err)
					os.Exit(1)
				}
				stringFlag = string(data)
			}

			if debug {
				fmt.Printf("Integer: %d => %s \n", intFlag, nconst.Name(intFlag))
				fmt.Printf("String: %s\n", stringFlag)
				fmt.Printf("Date: %s\n", dateFlag)
			}

			if err := scramble.Scramble(intFlag, stringFlag, dateFlag, debug); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().IntVarP(&intFlag, "constant", "i", 0, "An integer flag (0-5)")
	rootCmd.Flags().StringVarP(&stringFlag, "string", "s", "default", "A string flag")
	rootCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "A filepath to read string input from")
	rootCmd.Flags().BoolVarP(&debug, "debug", "v", false, "Enable verbose output")
	// 3 January 2009
	rootCmd.Flags().StringVarP(&dateFlag, "date", "d", "01-03-2009", "A date flag (MM-DD-YYYY)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
