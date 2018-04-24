package main


import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
//Man må åpne terminalen og skrive "go get -u github.com/gorilla/mux" for at det skal fungere
func main() {


	r := mux.NewRouter()

	r.HandleFunc("/1", response1)
	r.HandleFunc("/2", response2)
	r.HandleFunc("/3", response3)
	r.HandleFunc("/4", response4)
	r.HandleFunc("/5", response5)

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}


type County struct {
	Entries []struct {
		Name   string `json:"navn"`
		Number string `json:"nummer"`
	} `json:"entries"`
}

type Station struct {
	Entries []struct {
		Latitude  string `json:"latitude"`
		Name      string `json:"navn"`
		Plastic   string `json:"plast"`
		GlasMetal string `json:"glass_metall"`
		Shoe      string `json:"tekstil_sko,omitempty"`
		Longitude string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

type Job struct {
	Entries []struct {
		Yrke        string `json:"yrke"`
		Arbeidssted string `json:"arbeidssted_kommune"`
	} `json:"entries"`
}

type Sats struct {
	Entries []struct {
		Verdensdel  string `json:"verdensdel"`
		Kost		string `json:"kost"`
		Land	    string `json:"land"`
	} `json:"entries"`
}

type Video struct {
	Entries []struct {
		Title      string `json:"tittel"`
		Year		string `json:"aar"`
		Link	    string `json:"video_url"`
	} `json:"entries"`
}

var station Station
var county County
var job Job
var sats Sats
var video Video

func response1(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/difi/geo/fylke"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &county)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("counties.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, county)
	if err != nil {
		log.Fatal(err)
	}
}

func response2(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &station)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("stations.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, station)
	if err != nil {
		log.Fatal(err)

	}

}

func response3(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/nav/ledige-stillinger/2017?"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &job)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("jobs.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, job)
	if err != nil {
		log.Fatal(err)

	}

}

func response4(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/fad/reise/utland?"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sats)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("satser.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, sats)
	if err != nil {
		log.Fatal(err)
	}
}

func response5(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/nrk/norge-rundt?"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &video)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("videos.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, video)
	if err != nil {
		log.Fatal(err)
	}
}