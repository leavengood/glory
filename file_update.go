package glory

import (
	"fmt"
	"os"

	"github.com/kardianos/osext"
)

func updateExecutable(url, sha1 string) error {
	ex, err := osext.Executable()
	if err != nil {
		return err
	}

	newName := fmt.Sprintf("%s_new", ex)
	newFile, err := os.Create(newName)
	if err != nil {
		return err
	}

	err = downloadFile(url, newFile)
	if err != nil {
		return err
	}

	// Make sure file is saved to disk and seek back to the start
	newFile.Sync()
	newFile.Seek(0, 0)

	// Verify the checksum
	err = verifyChecksum(newFile, sha1)
	if err != nil {
		fmt.Println("Checksum does not match file")
		return err
	}

	// Make it executable
	newFile.Chmod(0755)

	newFile.Close()

	// Move current executable to <name>_old
	os.Rename(ex, fmt.Sprintf("%s_old", ex))
	// Move new executable in place
	os.Rename(newName, ex)

	return nil
}
