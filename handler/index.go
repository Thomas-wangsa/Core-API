package handler

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: for handle logic in each endpoint
|
*/

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))

}
