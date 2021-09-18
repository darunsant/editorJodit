package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
type Editor struct {
	ID 		primitive.ObjectID  `bson:"_id" json:"id,omitempty"`
    Name 	string				`bson:"name"`
    Form  	string				`bson:"form"`
}

type FormEditor struct{
	Name 	string				`bson:"name"`
    Form  	string				`bson:"form"`
}

const (
	connectionURI = "mongodb://127.0.0.1:27017"
	portWebServe = ":3000"
)

func GetMongoDbConnection() (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal("Can no connect Database", err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("editor").Collection("form")
	return collection,nil
}
func addForm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	client, _ := GetMongoDbConnection()

	var form FormEditor
	_ = json.NewDecoder(r.Body).Decode(&form)
	cur, _ := client.InsertOne(context.TODO(), form)
	fmt.Println(cur)
	json.NewEncoder(w).Encode(cur)

}
func getForm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var forms []Editor
	client, _ := GetMongoDbConnection()
	cur, _ := client.Find(context.TODO(), bson.M{})
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var form Editor
		err := cur.Decode(&form)
		if err != nil {
			log.Fatal(err)
		}
		forms = append(forms, form)
	}
	fmt.Println(forms)
	
	json.NewEncoder(w).Encode(forms)
}

func getFormById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var form Editor
	client, _ := GetMongoDbConnection()

	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := client.FindOne(context.TODO(), filter).Decode(&form)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(form)

}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/add/form", addForm).Methods("POST")
	r.HandleFunc("/api/get/form", getForm).Methods("GET")
	r.HandleFunc("/api/get/form/{id}",getFormById).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})


	log.Fatal(http.ListenAndServe(":8000",  handlers.CORS(headers, methods, origins)(r)))
}