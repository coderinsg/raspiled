package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"
	"bytes"
	"strconv"
)

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from hello handler!")
}

func updateDelay(w http.ResponseWriter, r *http.Request) {
	// ledDelay := strconv.Itoa(r.FormValue("blinkdelay"))
	blinkdelay, _ := strconv.ParseFloat(r.FormValue("blinkdelay"),32)
	blinkdelay = blinkdelay / 1000
	ledDelay := fmt.Sprintf("%f", blinkdelay)
	fmt.Println("ledDelay is", ledDelay)

	var out bytes.Buffer
	// cmd := exec.Command("date")
	// patchString := `'{"spec": {"template":{"spec":{"containers":[{"env":[{"name":"LED_DELAY","value":"` + ledDelay + `"}],"name":"raspiled"}]}}}}'`
	// patchString := `{"spec": {"template":{"spec":{"containers":[{"env":[{"name":"LED_DELAY","value": "1"}],"name":"raspiled"}]}}}}`
	patchString := `{"spec": {"template":{"spec":{"containers":[{"env":[{"name":"LED_DELAY","value":"` + ledDelay + `"}],"name":"raspiled"}]}}}}`

	fmt.Println("patchString is", patchString)

	cmd := exec.Command("kubectl","patch","deploy","raspiled","-p",patchString)
	cmd.Stdout = &out
	cmd.Run()
	// fmt.Fprintf(w, "Updated already!")
	fmt.Fprintf(w, out.String())
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/assets", updateDelay).Methods("POST")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}
