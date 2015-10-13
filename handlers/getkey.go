package handlers

import (
	"fmt"
	"net/http"
	"github.com/mukulrawat1986/recursepp/db"
)

func GetKey(d db.DB) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		key := r.URL.Query().Get("key")
		
		if key == ""{
			http.Error(w, "missing key name", http.StatusBadRequest)
			return
		}

		val, err := d.Get(key)
		if err == db.ErrNotFound{
			http.Error(w, "not found", http.StatusNotFound)
			return
		}else if err != nil{
			http.Error(w, fmt.Sprintf("error getting from database %s", err), http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))

	})
}
