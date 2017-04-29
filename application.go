package main


import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]Users)  //connected clients
var curUsers = make(chan Users) //holds thes] users


//upgrader taking a http socket and turnign it into websocket
var upgrader = websocket.Upgrader{}

//user struct

type Users struct {
	UUId string `json:uuid"`
	User string `json:user`
	longitude float64 `json:"longitude"`
	latitude float64 `json:"longitude"`
}

func main() {



	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	//configure websocket route
	router.HandleFunc("/ws", HandleConnections)
	router.HandleFunc("/users", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	//listening for incoming users
	go handleUsers()
	log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func HandleConnections(w http.ResponseWriter, r *http.Request){

	//upgrade initial get request to a  websocket
	ws, err := upgrader.Upgrade(w,r,nil)
	if err!= nil{
		log.Fatal(err)
	}

	//close connection when it returns
	defer ws.Close()




		var user Users
		//reads in a new user as ajson and maps it to a user object

		readErr := ws.ReadJSON(&user)
		if readErr != nil {
			log.Printf("error %v", err)
			delete(clients,ws)


		}
		//updates the websocket the client and maps to a user
		clients[ws] = user
		//sends user to list of current users
		curUsers <- user

}

func handleUsers(){
	for{
		//grabs the next user from broadcast channel??
		u := <- curUsers


		for client:= range clients{
			//if distance is less tahn 

		}

	}
}



func TodoIndex(w http.ResponseWriter, r *http.Request) {
	curUser := &Users{}
	curUser.UUId = r.FormValue("uuid")
	curUser.User = r.FormValue("name")
	curUser.longitude = r.FormValue("longitude")
	curUser.latitude = r.FormValue("latitude")
	fmt.Fprintln("we have received information")

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
