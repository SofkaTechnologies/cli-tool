package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type Updater interface {
	Get(destination string, url string) error
	Head(url string) (*http.Response, error)
}

func AutoUpdate(updater Updater) error {
	fmt.Println("Checking available versions of the cli for " + runtime.GOOS + "...")
	version := "M_h7YM2gafrSBWfkc1yLNo8wzRyU4cnA"
	binaryUrl := "https://cli-tool.s3.amazonaws.com/bin/" + runtime.GOOS + "/sofka-cli.zip"
	res, err := updater.Head(binaryUrl)
	if err != nil {
		if err != nil {
			return errors.New("error getting version info" + err.Error())
		}
	}
	const versionHeader = "X-Amz-Version-Id"
	if version != res.Header[versionHeader][0] {
		fmt.Println("Updating...")
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		destination := dir + "/updated"
		err = updater.Get(destination, binaryUrl)
		if err != nil {
			return errors.New("error downloading file" + err.Error())
		}
		fmt.Println("Done...")
		fmt.Println("Please rerun the program to use the new version downloaded to " + destination)
		os.Exit(1)
	} else {
		fmt.Println("Using latest  version")
	}
	return nil
}
