package handlers

import (
	"net/http"
	"github.com/mukulrawat1986/recursepp/db"
)

func SetKey(d db.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		values := r.URL.Query()

		var key, val string

		for k := range values{
			key = k
		}

		val = values[key][0]
		
		if key == ""{
			http.Error(w, "missing key name", http.StatusBadRequest)
			return
		}

		if err := d.Set(key, val); err != nil{
			http.Error(w, "error setting values in db", http.StatusInternalServerError)
			return
		}

		//fmt.Printf("%v\n", val)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("saved"))

	})
}