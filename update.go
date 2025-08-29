package main
import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update the application to the latest version",
		Long:  `This command updates the application to the latest available version.`,
		Run:   runUpdate,
	}
	forceUpdate bool
)

func init() {
	updateCmd.Flags().BoolVarP(&forceUpdate, "force", "f", false, "Force update even if already on the latest version")
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) {
	currentVersion := "1.0.0" // This would be dynamically determined in a real application
	latestVersion := getLatestVersion()

	if currentVersion == latestVersion && !forceUpdate {
		fmt.Println("You are already using the latest version:", currentVersion)