package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "../plugins/windows"
)

var people []Person

type Person struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    Address *Address `json:"address,omitempty"`
}

type Address struct {
    Country string `json:"country,omitempty"`
    City string `json:"city,omitempty"`
}

func DefaultHandler(w http.ResponseWriter, r *http.Request)  {    
    fmt.Println("DefaultHandler")
    fmt.Fprint(w, "Home")
}


func PeopleHandler(w http.ResponseWriter, r *http.Request){
    fmt.Println("PeopleHandler")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(people)
}

func PersonHandler(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("PersonHandler")
    w.Header().Set("Access-Control-Allow-Origin", "*")
     params := mux.Vars(r)
     for _,item := range people{
         if item.ID == params["id"]{
            json.NewEncoder(w).Encode(item)
         }
     }
}

func SystemInfoHndlr(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("SystemInfoHndlr")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    sysinfo := windows.Property()
    json.NewEncoder(w).Encode(sysinfo)
}


func main()  {
    fmt.Println("Server is running...");
    route := mux.NewRouter()
    people = append(people, Person{ID : "1", Name : "Morteza", Address : &Address{Country : "Iran", City : "Mashhad"}})
    people = append(people, Person{ID : "2", Name : "Navid", Address : &Address{Country : "Spain", City : "Madrd"}})
    people = append(people, Person{ID : "3", Name : "Hossein", Address : &Address{Country : "Monaco", City : "Mona"}})
    route.HandleFunc("/", DefaultHandler).Methods("Get")
    route.HandleFunc("/people", PeopleHandler).Methods("Get")
    route.HandleFunc("/people/{id}", PersonHandler).Methods("Get")
    
    route.HandleFunc("/fm-remove/{filename}", PersonHandler).Methods("Get")
    route.HandleFunc("/sysinfo", SystemInfoHndlr).Methods("Get")

    log.Fatal(http.ListenAndServe(":9000", route))
}
