<<<<<<< Tabnine <<<<<<<
package routes

import (
	"github.com/gorilla/mux"//-
	"github.com/adamsbite/flem-project-go-backend/pkg/controllers"//-
    "github.com/gorilla/mux"//+
    "github.com/adamsbite/flem-project-go-backend/pkg/controllers"//+
)//+

//-
var RegisterBookStoreRoutes = func(router *mux.Router)//-
//-
{	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")//-
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")//-
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")//-
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")//-
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")//-
var RegisterBookStoreRoutes = func(router *mux.Router) {//+
    router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")//+
    router.HandleFunc("/book/", controllers.GetBook).Methods("GET")//+
    router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")//+
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")//+
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")//+
}
>>>>>>> Tabnine >>>>>>>// {"conversationId":"3d619694-454a-4829-9748-653829bfa6f1","source":"instruct"}