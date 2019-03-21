package main

import (
	"net/http"
)

func List(aggr Aggregator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//convert Values map[string][]string to map[string]string
		q := map[string]string{}
		for k, v := range r.URL.Query() {
			q[k] = v[0]
		}

		result, timeout, err := aggr.Get(q)
		if timeout {
			http.Error(w, "Timeout Exceed", 408)
			return
		}
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		w.Write(result)
	})
}
