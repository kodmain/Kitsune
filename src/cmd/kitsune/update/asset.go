package update

import (
	"fmt"
	"io"
	"net/http"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/kodmain/kitsune/src/internal/storages/fs"
)

type asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func (a *asset) Download(destination string) error {
	aNameSplit := strings.SplitN(a.Name, "-", 2)
	binaryPath := filepath.Join(destination, aNameSplit[0])

	wheel, err := user.LookupGroup("wheel")
	if err != nil {
		return err
	}

	root, err := user.Lookup("root")
	if err != nil {
		return err
	}

	resp, err := http.Get(a.BrowserDownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := fs.CreateFile(binaryPath, &fs.CreateOption{
		User:  root,
		Group: wheel,
		Perms: 0755,
	})

	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err == nil {
		fmt.Println("Downloaded service:", aNameSplit[0])
	}

	return err
}
