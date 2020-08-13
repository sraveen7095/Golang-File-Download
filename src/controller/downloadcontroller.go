package controller

import(
	"net/http"
	"fmt"
	"io"
	"path"
"net/url"	
)
func Index(w http.ResponseWriter, r *http.Request) {
	fileUrl := "http://africau.edu/images/default/sample.pdf"

	filename, err := GetFilename(fileUrl)
	fmt.Println(filename)
	if err != nil {
		http.Error(w, "error url", 502)
		return
	}
	resp, err := http.Get(fileUrl)
	if err != nil {
		http.Error(w, "error url", 502)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Remote server error", 503)
		return
	}
	return
}

func GetFilename(inputUrl string) (string, error) {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}
	u.RawQuery = ""
	return path.Base(u.String()), nil
}

