package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func DownloadBeatmapbyName(BeatmapName string) string {
	uri := url.URL{Host: "osu.ppy.sh", Path: fmt.Sprintf("osu/%s", BeatmapName)}
	h, err := http.Get(uri.String())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer h.Body.Close()

	md5sum := md5.New()
	if _, err := io.Copy(md5sum, h.Body); err != nil {
		fmt.Println(err)
		return ""
	}

	f, err := os.OpenFile(fmt.Sprintf("data/map/%x", md5sum.Sum(nil)), os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()

	if _, err := io.Copy(f, h.Body); err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("data/map/%s", md5sum.Sum(nil))
}
