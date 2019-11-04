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
)

type Items struct {
	Items []Item		`json:"items"`
}

type Item struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"vanillaDescription"`
}

var items Items;

func getOneItem(w http.ResponseWriter, r *http.Request) {
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

func getAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items);
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!");
}

func main() {
	fmt.Println("Initializing")

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
	router := mux.NewRouter().StrictSlash(true);
	router.HandleFunc("/", homeLink);
	router.HandleFunc("/items", getAllItems).Methods("GET");
	router.HandleFunc("/items/{id}", getOneItem).Methods("GET");
	log.Fatal(http.ListenAndServe(":8080", router));
}