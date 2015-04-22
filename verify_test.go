package glory

import (
	"bytes"
	"testing"
)

const (
	Quote    = "We hold these truths to be self-evident: that all men are created equal; that they are endowed by their Creator with certain unalienable rights; that among these are life, liberty, and the pursuit of happiness."
	Checksum = "2120db300d88bf770baa835e268b53a3da04324f"
)

func TestVerifyChecksumSuccess(t *testing.T) {
	r := bytes.NewBufferString(Quote)
	err := VerifyChecksum(r, Checksum)
	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}
}

func TestVerifyChecksumFailure(t *testing.T) {
	r := bytes.NewBufferString("Not Jefferson")
	err := VerifyChecksum(r, Checksum)
	if err != InvalidChecksum {
		t.Errorf("Expected %v but received error: %v", InvalidChecksum, err)
	}
}
