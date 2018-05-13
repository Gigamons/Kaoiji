package server

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"git.gigamons.de/Gigamons/Kaoiji/tools/usertools"
	"github.com/google/uuid"

	"git.gigamons.de/Gigamons/Kaoiji/handlers"
	"git.gigamons.de/Gigamons/Kaoiji/objects"
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
		return
	} else if r.Header.Get("osu-token") != "" && objects.TokenExists(r.Header.Get("osu-token")) {
		handlers.HandlePackets(w, r, objects.GetToken(r.Header.Get("osu-token")))
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

	objects.StartTimeoutChecker()
	objects.NewToken(uuid.UUID{}, 0, 0, *usertools.GetUser(100))
	fmt.Printf("Kaoiji is listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
