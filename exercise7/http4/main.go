package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var itemList = template.Must(template.New("itemlist").Parse(`
<h1>item list</h1>
<table>
<tr style='text-align: left'>
  <th>name</th>
  <th>price</th>
</tr>
{{range $key, $val := .}}
<tr>
  <td>{{$key}}</td>
  <td>{{$val}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemList.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %q\n", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error: 'item' key does not exist.\n")
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: [%s] item already exists.\n", item)
		return
	}

	if price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error: 'price' key does not exist.\n")
		return
	}

	num, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	db[item] = dollars(num)
	fmt.Fprintf(w, "create success %s: %s\n", item, db[item])
}

// 商品価格を更新する。商品が無いか、値が不正であればエラーを報告する。
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		if price == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid price: %q\n", price)
			return
		}

		num, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid price: %q\n", price)
			return
		}

		db[item] = dollars(num)
		fmt.Fprintf(w, "update success %s: %s\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "delete success item: %q\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
