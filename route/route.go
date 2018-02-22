package route

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: routing list with interface method
|
|
*/

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	General "BDA_API/struct/general"
	V1 		"BDA_API/endpoint/v1"
)

type Router struct {
	mux *chi.Mux
}

/*
|--------------------------------------------------------------------------
| Function Information
|--------------------------------------------------------------------------
| @author 		: Thomas
| @Init_Route() : initializing api with go chi service
| @endpoint()	: set endpoint of BDA-API
| 				  call of API V1 package to initializing
|
*/

func (r *Router) Init_Route()  {
	r.mux = chi.NewRouter()
	r.end_point()
}

func (r *Router) end_point() {
	var v1 = &V1.V1_API{}
	v1.Engine = r.mux
	v1.Init_Endpoint()
}

/*
|--------------------------------------------------------------------------
| Function Information
|--------------------------------------------------------------------------
| @author 		: Thomas
| @Run() 		: Run the BDA-API with detail credentials
|
*/

func (r *Router) Run(app General.App) {
	log.Println(app.Name + " is Runing Now.")
	log.Println("Server " + app.Environment + " = " + app.Host + ":" + app.Port )
	http.ListenAndServe(":"+app.Port, r.mux)
}
