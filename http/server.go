package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DamonWright101/go-postgres-integration/logic"
	"github.com/golang/glog"
)

func Server(listener int) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
		fmt.Fprintf(w, BuildPersonTable([]logic.Person{}))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":"+strconv.Itoa(listener), nil)
	if err != nil {
		log.Fatal(err)
	}
	glog.Info("started listening on port" + strconv.Itoa(listener))
}

func BuildPersonTable(data []logic.Person) string {
	var table string
	table = `<table style="width:100%"><tr>` +
		`<th>ID</th>` +
		`<th>EMAIL</th>` +
		`<th>USERNAME</th>` +
		`<th>FIRST NAME</th>` +
		`<th>LAST NAME</th>` +
		`</tr>`

	table += `<tr></tr></td>`

	return table
}
