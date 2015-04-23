package glory

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

var invalidChecksum = errors.New("Invalid checksum")

// Verify that the given reader has the SHA1 checksum provided
func verifyChecksum(r io.Reader, checksum string) error {
	if generateChecksum(r) != checksum {
		return invalidChecksum
	}

	return nil
}

func generateChecksum(r io.Reader) string {
	h := sha1.New()

	_, err := io.Copy(h, r)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
