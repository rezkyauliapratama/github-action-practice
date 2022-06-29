package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	// DB     *sql.DB
}

type Apify struct {
	Ip string `json:"ip"`
}

// host, port, user, password, dbname string
func (a *App) Initialize() {

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) checkPublicIp(w http.ResponseWriter, r *http.Request) {
	var data Apify

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": data.Ip})
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.checkPublicIp).Methods("GET")

}
