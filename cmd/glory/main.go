package main

import (
	"fmt"
	"os"

	"github.com/leavengood/glory"
)

func Notify(notifyUrl, fileUrl, sha1, secret string) {
	ur := glory.NewUpdateRequest(fileUrl, sha1)
	ur.SendingNow()
	ur.Post(notifyUrl, secret)
}

func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s <server update URL> <update file URL> <update file SHA1> <shared secret>\n", os.Args[0])
		os.Exit(1)
	}

	Notify(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
}
