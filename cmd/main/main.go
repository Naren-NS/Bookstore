// package main

// import (
// 	"log"
// 	"net/http"

	
// 	"github.com/gorilla/mux"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/naren/go-bookstore/pkg/routes"
// )

// func main(){
// 	r := mux.NewRouter()
// 	routes.RegisterBookStoreRoutes(r)
// 	http.Handle("/",r)
// 	log.Fatal(http.ListenAndServe("localhost:9010", r))
// }

//tells that bookstore.routes has our routes
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/naren/go-bookstore/pkg/routes"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterBookStoreRoutes(r)

    // Serve static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("D:/go-bookstore/static"))))

    log.Fatal(http.ListenAndServe("localhost:9010", r))
}
