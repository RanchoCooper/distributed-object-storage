package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
	f, err := os.Open(os.Getenv(RootPath) + DefaultPrefix + objectName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()

	_, _ = io.Copy(w, f)
}
