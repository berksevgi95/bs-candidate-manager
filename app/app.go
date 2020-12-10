// package app

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jinzhu/gorm"
// 	"./controllers"
// 	"../config"
// 	"../migrate"
// )

// type App struct {
// 	Router *mux.Router
// 	DB     *gorm.DB
// }

// func (a *App) Initialize(config *config.Config) {
// 	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
// 		config.DB.Username,
// 		config.DB.Password,
// 		config.DB.Name,
// 		config.DB.Charset)

// 	db, err := gorm.Open(config.DB.Dialect, dbURI)
// 	if err != nil {
// 		log.Fatal("Could not connect database")
// 	}

// 	a.DB = migrate.DBMigrate(db)
// 	a.Router = mux.NewRouter()
// 	a.setRouters()
// }

// func (a *App) setRouters() {
// 	a.Post("/book", a.InputBook)
// 	a.Get("/book", a.ListBook)
// 	a.Get("/book/{codes:[1-9]+}", a.OneBook)
// 	a.Put("/book/{codes:[1-9]+}", a.UpdateBook)
// 	a.Delete("/book/{codes:[1-9]+}", a.DeletedBook)
// }

// //handler method
// func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("POST")
// }

// func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("Get")
// }

// func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("Put")
// }

// func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("Delete")
// }

// //get data
// func (a *App) InputBook(w http.ResponseWriter, r *http.Request) {
// 	controllers.InputBook(a.DB, w, r)
// }

// func (a *App) ListBook(w http.ResponseWriter, r *http.Request) {
// 	controllers.ListBook(a.DB, w, r)
// }

// func (a *App) OneBook(w http.ResponseWriter, r *http.Request) {
// 	controllers.OneBook(a.DB, w, r)
// }

// func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	controllers.UpdateBook(a.DB, w, r)
// }

// func (a *App) DeletedBook(w http.ResponseWriter, r *http.Request) {
// 	controllers.DeletedBook(a.DB, w, r)
// }

// //run
// func (a *App) Run(host string) {
// 	log.Fatal(http.ListenAndServe(host, a.Router))
// }

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"../config"
	"./models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router     *mux.Router
	Collection *mongo.Collection
}

func (a *App) Initialize(config *config.Config) {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	collection, _ := mongo.Connect(ctx, clientOptions)
	a.Collection = collection.Database("test").Collection("posts")
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/book", a.GetPeopleEndpoint)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Get")
}

func (a *App) GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []models.Post
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := a.Collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person models.Post
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

//run
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
