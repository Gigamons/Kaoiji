package server

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/Gigamons/common/logger"
	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/Kaoiji/handlers"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("cho-protocol", "19")
	header.Set("Connection", "keep-alive")
	header.Set("Keep-Alive", "timeout=5, max=100")
	header.Set("Content-Type", "text/html; charset=UTF-8")

	if r.Header.Get("osu-token") == "" && r.Header.Get("User-Agent") == "osu!" {
		handlers.LoginHandler(w, r)
		return
	} else if r.Header.Get("osu-token") != "" && !objects.TokenExists(r.Header.Get("osu-token")) {
		w.WriteHeader(403)
		fmt.Fprint(w, "Nya~")
		logger.Infof("Token %s got an Disconnect! token not found.", r.Header.Get("osu-token"))
		return
	} else if r.Header.Get("osu-token") != "" && objects.TokenExists(r.Header.Get("osu-token")) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			t := objects.GetToken(r.Header.Get("osu-token"))
			w.Write(t.Read())
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Errorln(err)
				return
			}
			handlers.HandlePackets(b, t)
		} else {
			flush, ok := w.(http.Flusher)
			if !ok {
				t := objects.GetToken(r.Header.Get("osu-token"))
				w.Write(t.Read())
				b, err := ioutil.ReadAll(r.Body)
				if err != nil {
					logger.Errorln(err)
					return
				}
				handlers.HandlePackets(b, t)
				return
			}
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			t := objects.GetToken(r.Header.Get("osu-token"))
			gz.Write(t.Read())
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Errorln(err)
				return
			}
			handlers.HandlePackets(b, t)
			flush.Flush()
		}
	}
}

func errHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("------------ERROR------------")
				fmt.Println(err)
				fmt.Println("---------ERROR TRACE---------")
				fmt.Println(string(debug.Stack()))
				fmt.Println("----------END ERROR----------")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// StartServer starts our HTTP Server.
func StartServer(port int) {
	r := mux.NewRouter()
	r.Use(errHandler)
	r.HandleFunc("/", main)

	objects.NewToken(uuid.UUID{}, 0, 0, usertools.GetUser(100))
	logger.Infof("Kaoiji is listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
