package glory

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// Download the given URL to the given writer
func downloadFile(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Unzip while downloading if the file is gzipped
	var r io.Reader
	if strings.HasSuffix(url, ".gz") {
		z, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer z.Close()
		r = z
	} else {
		r = resp.Body
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return nil
}
