// File: main.go
package main

import (
	"context"
    "encoding/json"
    "fmt"
    "log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


// for returing json format to client

type Person struct {
	Dob       string `json:"dob"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Timestamp string `json:"timestamp"`
}

type Contact struct {
	UserIDOne string `json:"UserIdOne"`
	UserIDTwo string `json:"UserIdTwo"`
	Timestamp string `json:"timestamp"`
}





//function to handle /contact 

func contactCreate(w http.ResponseWriter, r *http.Request) {
    var c Contact

   
    err := json.NewDecoder(r.Body).Decode(&c)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	coll := client.Database("tracingapi").Collection("Contact")
	
	//	fmt.Println("Connected to MongoDB!")        checking connection with db
	
	
	
	
	  insertResult, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"UserIdOne",  c.UserIDOne},
			{"UserIdTwo",  c.UserIDTwo},
			{"timestamp",  c.Timestamp}})
	fmt.Println("Inserted a single document: in contact collection ", insertResult.InsertedID)






	
	data, _ := json.Marshal(c)
	fmt.Println(string(data))
		w.Write(data)
		
}












// function to handle /users request 

func personCreate(w http.ResponseWriter, r *http.Request) {
    var p Person
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Println("Connected to MongoDB!")  checking connection with db


coll := client.Database("tracingapi").Collection("Users")

	insertResult, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"id",p.ID},
			{"name", p.Name},
			{"dob", p.Dob},
			{"phone", p.Phone},
			{"email", p.Email},
			{"timestamp", p.Timestamp }})
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

		
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
		w.Write(data)
}






func main() {
	
	mux := http.NewServeMux()
	
	mux.HandleFunc("/users", personCreate)      // handling /users case 1.
 
	mux.HandleFunc("/contact", contactCreate)   // handling /contact case 1.

	ewrr := http.ListenAndServe(":8080",mux)    //using port 8080 localhost for creating web server 
    log.Fatal(ewrr)
}