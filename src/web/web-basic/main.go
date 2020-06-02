package main

import(
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name": "Salman Saputra",
			"Message": "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Alias nostrum deserunt autem, voluptatibus consequuntur? Sit quis non sed neque id beatae voluptatum! Dolorum earum repudiandae delectus? Odio, illum. Nisi, rerum.",
		}

		var t, err = template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "About Page")
	// })

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}