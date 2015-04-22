package glory

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

var InvalidChecksum = errors.New("Invalid checksum")

// Verify that the given reader has the SHA1 checksum provided
func VerifyChecksum(r io.Reader, checksum string) error {
	h := sha1.New()

	_, err := io.Copy(h, r)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", h.Sum(nil)) != checksum {
		return InvalidChecksum
	}

	return nil
}
