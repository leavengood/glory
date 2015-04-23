package glory

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const dateFormat = "20060102150405"

var (
	expiredTimeout = time.Second * 10
	updateError    = errors.New("Update error")
)

type UpdateRequest struct {
	url  string
	sha1 string
	time time.Time
}

func NewUpdateRequest(url, sha1 string) *UpdateRequest {
	return &UpdateRequest{url: url, sha1: sha1}
}

func (ur *UpdateRequest) SetTimestamp(timestamp string) {
	t, err := time.Parse(dateFormat, timestamp)
	if err == nil {
		ur.time = t
	}
}

func (ur *UpdateRequest) Timestamp() string {
	return ur.time.Format(dateFormat)
}

func (ur *UpdateRequest) SendingNow() {
	ur.time = time.Now().UTC()
}

func (ur *UpdateRequest) Expired(now time.Time) bool {
	return now.After(ur.time.Add(expiredTimeout))
}

func (ur *UpdateRequest) Signature(secret string) string {
	m := fmt.Sprintf("%s%s%s%s", ur.url, ur.sha1, ur.Timestamp(), secret)

	return generateChecksum(bytes.NewBufferString(m))
}

func (ur *UpdateRequest) Post(notifyUrl, secret string) error {
	resp, err := http.PostForm(notifyUrl,
		url.Values{"url": {ur.url}, "sha1": {ur.sha1}, "timestamp": {ur.Timestamp()},
			"signature": {ur.Signature(secret)}})

	if err != nil {
		return err
	}

	fmt.Println(resp)

	if resp.StatusCode != 200 {
		return updateError
	}

	return nil
}
