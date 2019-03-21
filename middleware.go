package main

import (
	"net/http"
	"regexp"
)

//Set all neccessary headers defined in main.go
func SetHeaders(h http.Handler, headers map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		h.ServeHTTP(w, r)
	})
}

//Allow only GET method
func GET(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			w.Write([]byte("Method not allowed"))
			return
		}
		h.ServeHTTP(w, r)
	})
}

//Validate query by matching regex pattern ?reg=54&&city=г.Новосибирск&&street=ул.Богдана Хмельницкого&&zip=630110
func QValid(h http.Handler, required []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pattern := map[string]string{
			"reg":    `^(\d{2})$`,
			"city":   `^(г\.[\-А-я]{4,27})$`,
			"street": `^([\-\.\\ \w]{3,55})$`,
			"zip":    `^([\d]{6})`,
		}
		q := r.URL.Query()
		for _, k := range required {
			match, _ := regexp.MatchString(pattern[k], q.Get(k))
			if !match {
				w.WriteHeader(401)
				w.Write([]byte("Invalid query"))
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}

/*
func exampleMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Our middleware logic goes here...
    next.ServeHTTP(w, r)
  })
}
*/
