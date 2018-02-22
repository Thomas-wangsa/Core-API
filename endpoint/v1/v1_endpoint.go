package v1_endpoint

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: lists all endpoint of BDA-API V1
| 			  import go chi as router service : https://github.com/go-chi/chi
| 			  import package handler to separated the process of each endpoint
|
*/

import (
	"github.com/go-chi/chi"
	Handler "BDA_API/handler"
)


type V1_API struct {
	Engine *chi.Mux
}

/*
|--------------------------------------------------------------------------
| Function Init_Endpoint
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: lists all endpoint of BDA-API V1
|
*/

func (v1 *V1_API)Init_Endpoint() {
	var r = v1.Engine
	r.Get("/", Handler.Index)

}