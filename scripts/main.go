package main

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func main() {

	pwd, err := os.Getwd()
	basePath := path.Dir(pwd)
	logDir := filepath.Join(basePath, "logs")

	execPath := "/usr/bin/terraform"
	tf, err := tfexec.NewTerraform(basePath, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	// Read env  variables
	OKTA_ORG_NAME := os.Getenv("OKTA_ORG_NAME")
	OKTA_BASE_URL := os.Getenv("OKTA_BASE_URL")
	OKTA_API_TOKEN := os.Getenv("OKTA_API_TOKEN")
	SELECT_WORKSPACE := os.Args[1]

	if SELECT_WORKSPACE == "" {
		log.Fatalf("expecting argument for selecting workspace: %s", err)
	}

	log.Printf("Selecting workspace:%s", SELECT_WORKSPACE)

	// Set env variables
	tf.SetEnv(map[string]string{"OKTA_ORG_NAME": OKTA_ORG_NAME, "OKTA_BASE_URL": OKTA_BASE_URL, "OKTA_API_TOKEN": OKTA_API_TOKEN})

	//setup logs
	tfLogPath := filepath.Join(logDir, "tf.log")

	err = tf.SetLog("INFO")

	if err != nil {
		log.Fatalf("unexpected SetLog error: %s", err)
	}

	err = tf.SetLogPath(tfLogPath)

	// select environment
	err = tf.WorkspaceSelect(context.Background(), SELECT_WORKSPACE)

	if err != nil {
		log.Fatalf("error selecting workspace: %s", err)
	}

	err = tf.Apply(context.Background())
	if err != nil {
		log.Fatalf("error running Apply: %s", err)
	}
	log.Printf("Done applying changes")
}
