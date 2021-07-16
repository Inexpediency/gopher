package web

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Element struct {
	Name    string
	Surname string
	Id      string
}

type Storage struct {
	Data     map[string]Element
	FilePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{
		Data:     make(map[string]Element),
		FilePath: filePath,
	}
}

func (s *Storage) save() error {
	log.Println("Saving", s.FilePath)

	err := os.Remove(s.FilePath)
	if err != nil {
		log.Println(err)
	}

	df, err := os.Create(s.FilePath)
	if err != nil {
		log.Println("Cannot create", s.FilePath)
		return err
	}
	defer func(df *os.File) {
		err := df.Close()
		if err != nil {

		}
	}(df)

	encoder := gob.NewEncoder(df)
	err = encoder.Encode(s.Data)
	if err != nil {
		log.Println("Cannot save to", s.FilePath)
		return err
	}

	return nil
}

func (s *Storage) load() error {
	log.Println("Loading", s.FilePath)

	loadFrom, err := os.Open(s.FilePath)
	defer func(loadFrom *os.File) {
		err := loadFrom.Close()
		if err != nil {

		}
	}(loadFrom)
	if err != nil {
		log.Println("Empty key/value storage!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	err = decoder.Decode(&s.Data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Print() {
	for k, e := range s.Data {
		fmt.Printf("key: %s value: %v\n", k, e)
	}
}

func (s *Storage) Add(k string, e Element) bool {
	if k == "" {
		return false
	}

	if s.LookUp(k) != nil {
		return false
	}

	s.Data[k] = e

	return true
}

func (s *Storage) Delete(k string) bool {
	if s.LookUp(k) == nil {
		return false
	}

	delete(s.Data, k)

	return true
}

func (s *Storage) Change(k string, e Element) bool {
	s.Data[k] = e

	return true
}

func (s *Storage) LookUp(k string) *Element {
	e, ok := s.Data[k]
	if ok {
		return &e
	}

	return nil
}

var storage *Storage

const insertTemplatePath = "./web/insert.gohtml"
const homeTemplatePath = "./web/home.gohtml"
const updateTemplatePath = "./web/update.gohtml"

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving", r.Host, "for", r.URL.Path)
	t := template.Must(template.ParseGlob(homeTemplatePath))
	err := t.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Println(err)

		return
	}
}

func listAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Listening the contents of the KV store")

	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">List all elements!</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\">Change an element!</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\">Insert new element!</a>")
	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")
	for k, v := range storage.Data {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	log.Println("Changing an element of the KV store!")

	tmpl := template.Must(template.ParseFiles(updateTemplatePath))
	if r.Method != http.MethodPost {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}

	key := r.FormValue("key")
	n := Element{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	if !storage.Change(key, n) {
		log.Println("Update operation failed!")

		return
	}

	err := storage.save()
	if err != nil {
		log.Println(err)

		return
	}

	err = tmpl.Execute(w, struct{ Success bool }{true})
	if err != nil {
		log.Println(err)

		return
	}
}

func insertElement(w http.ResponseWriter, r *http.Request) {
	log.Println("Inserting an element to the KV store!")
	tmpl := template.Must(template.ParseFiles(insertTemplatePath))
	if r.Method != http.MethodPost {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}
	key := r.FormValue("key")
	n := Element{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}
	if !storage.Add(key, n) {
		log.Println("Add operation failed!")

		return
	}

	err := storage.save()
	if err != nil {
		log.Println(err)

		return
	}

	err = tmpl.Execute(w, struct{ Success bool }{true})
	if err != nil {
		log.Println(err)

		return
	}
}

func StartKeyValueStorageServer(port string) {
	storage = NewStorage("/tmp/kvdata.gob")
	err := storage.load()
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
