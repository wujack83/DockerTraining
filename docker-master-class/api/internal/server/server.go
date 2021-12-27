package server

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"

	"gitlab.com/andersph/docker-master-class/api/internal"
	"gitlab.com/andersph/docker-master-class/api/internal/database"

)

func healthCheck(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "api is healthy")
    log.Println("health check endpoint hit")
}

func newEmployee(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request   
	reqBody, _ := ioutil.ReadAll(r.Body)
	
    var employee internal.Employee 
	json.Unmarshal(reqBody, &employee)

	// try 5 times to write the employee entry to database, wait 1 second between
	message := database.SqlWriterRetry(employee, 5, 1)
	log.Println(message)
	fmt.Fprintf(w, "%+v", message)
}

func HandleRequests(configFile string) {

	// load configuration
	c := internal.LoadConfig()
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/healthy", healthCheck)
	myRouter.HandleFunc(c.Endpoints.CreateEmployee, newEmployee).Methods("POST")
	log.Println("listen on port", internal.GetEnv("PORT", c.Server.Port))
    log.Fatal(http.ListenAndServe(":"+internal.GetEnv("PORT", c.Server.Port), myRouter))
}
