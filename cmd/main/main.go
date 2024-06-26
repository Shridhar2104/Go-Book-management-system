package main
import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	_  "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Shridhar2104/Go-Book-management-system/pkg/routes"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}



