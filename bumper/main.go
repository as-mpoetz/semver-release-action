package main

import (
	"os"

	"github.com/K-Phoen/semver-release-action/bumper/event"
	"github.com/K-Phoen/semver-release-action/bumper/git"
	"github.com/K-Phoen/semver-release-action/bumper/release"
	"github.com/K-Phoen/semver-release-action/bumper/semver"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "bumper",
	}

	rootCmd.AddCommand(semver.Command())
	rootCmd.AddCommand(release.Command())
	rootCmd.AddCommand(event.GuardCommand())
	rootCmd.AddCommand(event.IncrementCommand())
	rootCmd.AddCommand(git.LatestTagCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
