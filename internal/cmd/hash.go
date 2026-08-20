package cmd

import (
	"fmt"
	"github.com/exec-cmd/hasher/internal/filehash"

	"github.com/spf13/cobra"
)

var algorithm string

var hashCmd = &cobra.Command{
	Short: "Hash a file",
	Long:  "Hash a file using the specified algorithm",

	Use:          "hash <file>",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),

	RunE: runE,
}

func runE(cmd *cobra.Command, args []string) error {
	h, err := filehash.HashFromAlgorithm(filehash.Algorithm(algorithm))
	if err != nil {
		return fmt.Errorf(
			"invalid hash algorithm %q, available: %v: %w\n",
			algorithm,
			filehash.Algorithms,
			err,
		)
	}

	sum, err := filehash.HashFile(args[0], h)
	if err != nil {
		return err
	}

	fmt.Println(sum)
	return nil
}

func init() {
	hashCmd.Flags().StringVarP(&algorithm, "alg", "a", "sha256", "Hash algorithm")

	rootCmd.AddCommand(hashCmd)
}
