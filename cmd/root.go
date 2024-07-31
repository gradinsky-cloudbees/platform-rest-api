package cmd

import (
	rest "github.com/gradinsky-cloudbees/platform-rest-api/rest"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	cmd = &cobra.Command{
		Use:   "rest-api",
		Short: "CLI to execute REST API actions",
		Long:  "CLI to execute REST API actions",
		RunE:  run,
	}
	cfg rest.Config
)

// These are the input flags. Created with external/rest/types.go
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

func run(*cobra.Command, []string) error {
	resp, err := rest.ExecuteApiCall(cfg.RequestType, cfg.Url, cfg.Payload, cfg.BearerToken, cfg.Username, cfg.Password, cfg.ExpectedResponseCode)
	if err != nil {
		log.Println(err.Error())
		//This is to create an output that can be read by a future step. All unique outputs must be in their own files.
		err2 := os.WriteFile(filepath.Join(os.Getenv("CLOUDBEES_OUTPUTS"), "response"), []byte(err.Error()), 0666)
		if err2 != nil {
			return err2
		}
		os.Exit(1)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Println("Successful API call")
	//Write output when successful, it parses the response body
	err = os.WriteFile(filepath.Join(os.Getenv("CLOUDBEES_OUTPUTS"), "response"), []byte(bodyString), 0666)

	return err
}
