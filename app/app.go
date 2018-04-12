package app

import (
	"net/http"
	"dumper/mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
	"bytes"
)

func Ini(porta string) {
	http.HandleFunc("/", lista)
	http.HandleFunc("/dump", dump)
	http.HandleFunc("/get", get)
	http.ListenAndServe(porta, nil)
}

func lista(w http.ResponseWriter, r *http.Request) {
	err := filtraToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fis, err := ioutil.ReadDir("./static")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	data := make([]string, 0)
	for _, fi := range fis {
		data = append(data, fi.Name())
	}
	printJson(w, data)
}

func dump(w http.ResponseWriter, r *http.Request) {
	err := filtraToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	host := r.URL.Query().Get("host")
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")
	db := r.URL.Query().Get("db")
	if host == "" || user == "" || pass == "" || db == "" {
		http.Error(w, "host, user, pass e db obrigatórios", http.StatusBadRequest)
		return
	}
	file, err := mysql.Dump(host, user, pass, db, "./static")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Erro bool
		File string
	}{
		false,
		file,
	}
	printJson(w, data)
}

func get(w http.ResponseWriter, r *http.Request) {
	err := filtraToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	file := r.URL.Query().Get("file")
	data, err := ioutil.ReadFile(filepath.Join("./static", file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=" + file)
	http.ServeContent(w,r, filepath.Join("./static", file), time.Now(), bytes.NewReader(data))
}

func printJson(w http.ResponseWriter, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-type", "application/json; charset=utf8")
	fmt.Fprint(w, string(bytes))
	return nil
}
