package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
	f, err := os.Create(os.Getenv(RootPath) + DefaultPrefix + objectName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, _ = io.Copy(f, r.Body)
}