package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

type Welcome struct {
	Name string
	Time string
}

type JsonResponse struct {
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	JsonNested JsonNested `json:"JsonNested"`
}

type JsonNested struct {
	NestedValue1 string `json:"nestedKey1"`
	NestedValue2 string `json:"nestedKey2"`
}

type JsonInfo struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Address Address `json:"Address"`
	ContactInfo ContactInfo `json:"ContactInfo"`
}

type Address struct {
	Street string `json:"street"`
	City string `json:"city"`
}

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	nested := JsonNested{
		NestedValue1 : "first nested value",
		NestedValue2 : "second nested value",
	}
	jsonResp := JsonResponse{
		Value1: "some Data",
		Value2: "other Data",
		JsonNested: nested,
	}

	myAddress := Address {
		Street: "4225 University Avenue",
		City: "Columbus",
	}
	
	myContactInfo := ContactInfo {
		Email: "ThisIsMyEmail@gmail.com",
		Phone: "706-123-4567",
	}

	myJsonInfo := JsonInfo {
		Firstname: "Jay",
		Lastname:"Son",
		Address: myAddress,
		ContactInfo: myContactInfo,
	}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) 

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(jsonResp)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(myJsonInfo)
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}