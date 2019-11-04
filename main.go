package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type Items struct {
	Items []Item		`json:"items"`
}

type Item struct {
	ID					int		`json:"id"`
	Name				string	`json:"name"`
	Description			string	`json:"description"`
	Atk					int		`json:"atk"`
	matk                int     `json:"matk"`
	Def                 int     `json:"def"`
	Mdef                int     `json:"mdef"`
	Speed               int     `json:"speed"`
	Jump                int     `json:"jump"`
	Kb                  int     `json:"kb"`
	Aspd                int     `json:"aspd"`
	Cspd                int     `json:"cspd"`
	Crit                int     `json:"crit"`
	Cdmg                int     `json:"cdmg"`
	CompoundType        int     `json:"compoundType"`
	Cost                int     `json:"cost"`
	CurseId             int     `json:"curseId"`
	CurseMaxKills       int     `json:"curseMaxKills"`
	ForceRenderBelow    bool	`json:"forceRenderBelow"`
	GemRegionX          int     `json:"gemRegionX"`
	GemRegionY          int     `json:"gemRegionY"`
	Hotkey              int     `json:"hotkey"`
	Icon                int     `json:"icon"`
	EggId               int     `json:"eggId"`
	MiniGemRegion       int     `json:"miniGemRegion"`
	ShowHair            bool	`json:"showHair"`
	StabBonus           int     `json:"stabBonus"`
	Subtype             string  `json:"subtype"`
	Success             int     `json:"success"`
	Type                int     `json:"type"`
	VisualName          string  `json:"visualName"`
}

var items Items;

func showItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("error parsing")
		return
	} 
	for i := 0; i < len(items.Items); i++ {
		if items.Items[i].ID == id {
			json.NewEncoder(w).Encode(items.Items[i]);
		}
	}
}

func listItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items);
}

func main() {
	fmt.Println("Initializing...")

	// Attempt to open json file
	jsonFile, err := os.Open("items.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened items.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'items' which we defined above
	json.Unmarshal(byteValue, &items)

	// Create our router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/items", listItems).Methods("GET")
	router.HandleFunc("/v1/items/{id}", showItemById).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))

	// Listen on whatever port heroku wants. 8080 by default.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}