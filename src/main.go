package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

var volumeServer = flag.String("volumeServer", "http://localhost:18080", "")

func main() {
	flag.Parse()
	const bindAddress = ":80"
	http.HandleFunc("/", requestHandler)
	fmt.Println("Http server listening on", bindAddress)
	_ = http.ListenAndServe(bindAddress, nil)
}

func requestHandler(response http.ResponseWriter, request *http.Request) {
	if !strings.HasSuffix(request.URL.Path, ".webp") {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	path := request.URL.Path
	vid, fid, fileName, _, _ := parseURLPath(path)
	var imgUrl = *volumeServer + "/" + vid + "/" + fid + "/" + strings.TrimRight(fileName, ".webp")
	resp, err := http.Get(strings.TrimRight(imgUrl, "/"))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	segments := make([]string, 0)
	query := request.URL.Query()
	for q := range query {
		key := q
		element := request.FormValue(key)
		if element == "true" {
			segments = append(segments, fmt.Sprintf("-%v", key))
		} else {
			segments = append(segments, fmt.Sprintf("-%v", key), fmt.Sprintf("%v", element))
		}
	}
	segments = append(segments, "-o", "-", "--", "-")
	cmd := exec.Command("cwebp", segments...)
	cmd.Stdin = io.Reader(resp.Body)
	cmd.Stdout = response
	_ = cmd.Start()
	defer cmd.Wait()

	response.Header().Set("Content-Type", "image/webp")
	response.WriteHeader(http.StatusOK)
}

func parseURLPath(path string) (vid, fid, filename, ext string, isVolumeIdOnly bool) {
	switch strings.Count(path, "/") {
	case 3:
		parts := strings.Split(path, "/")
		vid, fid, filename = parts[1], parts[2], parts[3]
		ext = filepath.Ext(filename)
	case 2:
		parts := strings.Split(path, "/")
		vid, fid = parts[1], parts[2]
		dotIndex := strings.LastIndex(fid, ".")
		if dotIndex > 0 {
			ext = fid[dotIndex:]
			fid = fid[0:dotIndex]
		}
	default:
		sepIndex := strings.LastIndex(path, "/")
		commaIndex := strings.LastIndex(path[sepIndex:], ",")
		if commaIndex <= 0 {
			vid, isVolumeIdOnly = path[sepIndex+1:], true
			return
		}
		dotIndex := strings.LastIndex(path[sepIndex:], ".")
		vid = path[sepIndex+1 : commaIndex]
		fid = path[commaIndex+1:]
		ext = ""
		if dotIndex > 0 {
			fid = path[commaIndex+1 : dotIndex]
			ext = path[dotIndex:]
		}
	}
	return
}
