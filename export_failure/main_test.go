package main

import (
	"context"
	"io/ioutil"
	"testing"

	"dagger.io/dagger"
	log "github.com/sirupsen/logrus"
)

// TestExitZero runs a script that creates files, then exits 0. This is the happy path - we get our test files.
func TestExitZero(t *testing.T) {
	ctx := context.Background()
	scriptPath := "will_export.sh"
	outputDir := "output_success"
	client, err := dagger.Connect(ctx)
	if err != nil {
		t.Fatalf("error connecting to Dagger: %v", err)
	}
	defer client.Close()

	// get `ubuntu` image and execute the script from the given path
	img, err := client.Container().From("ubuntu:latest").WithExec([]string{"mkdir", "/tmp/test"}).
		WithFile("/tmp/test/script.sh", client.Host().File(scriptPath)).
		WithExec([]string{"/tmp/test/script.sh"}).
		Sync(ctx)
	if err != nil {
		log.Infof("error running script: %v", err)
	}

	_, err = img.Directory("/tmp/test").Export(ctx, outputDir)
	if err != nil {
		t.Fatalf("error exporting directory: %v", err)
	}

	files, err := ioutil.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("error reading output directory: %v", err)
	}

	if err != nil {
		t.Fatalf("error listing files: %v", err)
	}
	if len(files) == 0 {
		t.Fatalf("no files found in the output directory")
	}
}

// TestExitOne runs a script that creates files, then exits 1.
// This simulates something like gradle tests running, producing Junit ouput, and failing.
// In this scenario, I need the files generated by the failed container.
func TestExitOne(t *testing.T) {
	ctx := context.Background()
	scriptPath := "wont_export.sh"
	outputDir := "output_success"
	client, err := dagger.Connect(ctx)
	if err != nil {
		t.Fatalf("error connecting to Dagger: %v", err)
	}
	defer client.Close()

	// get `ubuntu` image and execute the script from the given path
	img, err := client.Container().From("ubuntu:latest").WithExec([]string{"mkdir", "/tmp/test"}).
		WithFile("/tmp/test/script.sh", client.Host().File(scriptPath)).
		WithExec([]string{"/tmp/test/script.sh"}).
		Sync(ctx)
	if err != nil {
		log.Infof("error running script: %v", err)
	}

	_, err = img.Directory("/tmp/test").Export(ctx, outputDir)
	if err != nil {
		t.Fatalf("error exporting directory: %v", err)
	}

	files, err := ioutil.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("error reading output directory: %v", err)
	}

	if err != nil {
		t.Fatalf("error listing files: %v", err)
	}
	if len(files) == 0 {
		t.Fatalf("no files found in the output directory")
	}
}
