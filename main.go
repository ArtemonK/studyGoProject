package main
import ("fmt"
	"net/http")

func home_page(w http.ResponseWriter, r *http.Request) {	//1 команда ответов 2 запрос
	fmt.Fprintf(w, "Go is suoer easy!")	//форматированная строка
	}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page) //отслеживает переход и выполняет функцию
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil) //порт который слушаем
}

func main() {
	handleRequest()
}