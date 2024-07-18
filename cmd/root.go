package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"rest-api/internal/rest-api"
)

var (
	cmd = &cobra.Command{
		Use:   "rest-api",
		Short: "CLI to execute REST API actions",
		Long:  "CLI to execute REST API actions",
		RunE:  run,
	}
	cfg rest_api.Config
)

func init() {
	cmd.Flags().StringVar(&cfg.Url, "url", "", "REST API URL")
	cmd.Flags().StringVar(&cfg.RequestType, "requestType", "", "Request Type [GET|POST|PUT|DELETE]")
	cmd.Flags().StringVar(&cfg.Payload, "payload", "", "Payload")
	cmd.Flags().StringVar(&cfg.BearerToken, "bearerToken", "", "Bearer token for authentication")
	cmd.Flags().StringVar(&cfg.Username, "username", "", "Username for basic authentication")
	cmd.Flags().StringVar(&cfg.Password, "password", "", "Password for basic authentication")
	cmd.Flags().StringVar(&cfg.ExpectedResponseCode, "expectedResponseCode", "", "Expected Response code [200|300|400|500]")
}

func Execute() error {
	return cmd.Execute()
}

func run(cmd *cobra.Command, args []string) error {
	resp, err := cfg.ExecuteApiCall(cfg)
	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println("-----\n", bodyString, err)
	//Write output
	err = os.WriteFile(filepath.Join(os.Getenv("CLOUDBEES_OUTPUTS"), "response"), []byte(bodyString), 0666)
	if err != nil {
		return err
	}

	return err
}
