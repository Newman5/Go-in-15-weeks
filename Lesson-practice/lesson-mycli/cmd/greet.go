package cmd
import (
 "fmt"
 "github.com/spf13/cobra"
 "strings" // add this one for strings.ToUpper
)

var (
    uppercase bool
    times     int
)

var greetCmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone by name",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		    if verbose {
        fmt.Println("Verbose mode enabled")
    }
        name := args[0]
        if uppercase {
            name = strings.ToUpper(name)
        }
        for i := 0; i < times; i++ {
            fmt.Printf("Hello, %s!\n", name)
        }
    },
}

func init() {
    rootCmd.AddCommand(greetCmd)
    greetCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "Print name in uppercase")
    greetCmd.Flags().IntVarP(&times, "times", "t", 1, "Number of times to greet")
}

