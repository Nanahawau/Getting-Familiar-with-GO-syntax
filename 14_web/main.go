package main

import ( "fmt" 
		"net/http" 
	)


	func index(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello Word")
	}

func main ()  {
http.HandleFunc("/", index);
fmt.Println("Server Starting")
http.ListenAndServe(":3000", nil)
}