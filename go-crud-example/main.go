package main

import (
	// inner dependencies :
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	// outer dependency :
	"github.com/gorilla/mux"
)

type Video struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var videos []Video

func getVideos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(videos)
}

func getVideo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params:= mux.Vars(r)
	for _,v :=range videos{
		if v.ID==params["id"]{
			json.NewEncoder(w).Encode(v)
			break
		}
	}
}

func deleteVideo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	params:=mux.Vars(r)

	for i,v:=range videos{
		if v.ID == params["id"]{
			videos=append(videos[:i],videos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(videos)
}

func createVideo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	var vid Video
	_ = json.NewDecoder(r.Body).Decode(&vid)
	vid.ID = strconv.Itoa(rand.Intn(1000000000))
	videos = append(videos,vid)
}

func updateVideo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params:= mux.Vars(r)
	for i:=range videos{
		if videos[i].ID==params["id"]{
			var updatedVid Video
			_ = json.NewDecoder(r.Body).Decode(&updatedVid)
			videos[i] = updatedVid
			json.NewEncoder(w).Encode(updatedVid)
			return
		}
	}
}

func main(){
	fmt.Println("fuck")

	r:=mux.NewRouter()

	videos = append(videos,Video{"1","5553535","First Video",&Author{"John","Doe"}})
	videos = append(videos,Video{"2","5553636","Second Video",&Author{"Jahne","Doe"}})

	r.HandleFunc("/videos",getVideos).Methods("GET")
	r.HandleFunc("/videos/{id}",getVideo).Methods("GET")
	r.HandleFunc("/videos",createVideo).Methods("POST")
	r.HandleFunc("/videos/{id}",updateVideo).Methods("PUT")
	r.HandleFunc("/videos/{id}",deleteVideo).Methods("DELETE")

	defer fmt.Println("Started server at :8080")

	log.Fatal(http.ListenAndServe(":8080",r))
}
