package utils

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

type updaterVersionInfoMock struct {
}

func TestAutoUpdateVersionInfo(t *testing.T) {
	var mock updaterVersionInfoMock
	err := AutoUpdate(mock)
	if err == nil || !strings.Contains(err.Error(), "error getting version info") {
		t.Fatal("should return 'error getting version info'")
	}
}

func (updaterVersionInfoMock) Get(destination string, url string) error {
	return nil
}

func (updaterVersionInfoMock) Head(url string) (*http.Response, error) {
	return nil, errors.New("FAILED")
}

type updaterDownloadMock struct {
}

func TestAutoUpdateDownload(t *testing.T) {
	var mock updaterDownloadMock
	err := AutoUpdate(mock)
	if err == nil || !strings.Contains(err.Error(), "error downloading file") {
		t.Fatal("should return 'error downloading file'")
	}
}

func (updaterDownloadMock) Head(url string) (*http.Response, error) {
	return &http.Response{Header: map[string][]string{"X-Amz-Version-Id": {"111111"}}}, nil
}

func (updaterDownloadMock) Get(destination string, url string) error {
	return errors.New("FAILED")
}
