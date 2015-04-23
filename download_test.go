package glory

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func downloadFileTestHelper(t *testing.T, urlSuffix, served, expected string) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, served)
	}))
	defer ts.Close()

	b := new(bytes.Buffer)
	err := downloadFile(fmt.Sprintf("%s/%s", ts.URL, urlSuffix), b)
	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}
	result := b.String()
	if result != expected {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

const FileContents = "An updated file"

func TestDownloadFile_NormalFile(t *testing.T) {
	downloadFileTestHelper(t, "", FileContents, FileContents)
}

func TestDownloadFile_GzippedFile(t *testing.T) {
	expected := FileContents

	// Prepare zipped "file"
	z := new(bytes.Buffer)
	gz := gzip.NewWriter(z)
	gz.Write([]byte(expected))
	gz.Close()
	served := string(z.Bytes())

	downloadFileTestHelper(t, "file.gz", served, expected)
}
