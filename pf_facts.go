package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Rule struct {
	Id   int64  `json:"_id"`
	Text string `json:"text"`
}

type Img struct {
	//background images coming
}

type Tag struct {
	RuleId int64
	ImgId  int64
	Text   string
}

var templates = template.Must(template.ParseFiles("templates/rule.html"))

func renderTemplate(w http.ResponseWriter, template string, rule *Rule) {
	if err := templates.ExecuteTemplate(w, template, rule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadRule(ruleId string) (rule *Rule, err error) {
	rule = &Rule{}
	var ruleBytes []byte
	filename := "./data/" + ruleId + ".json"
	ruleBytes, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(ruleBytes, rule)
	if err != nil {
		return
	}
	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ruleId := strconv.Itoa(rand.Intn(2) + 1)
	http.Redirect(w, r, "/rule/"+ruleId, http.StatusFound)
}

func ruleHandler(w http.ResponseWriter, r *http.Request) {
	ruleId := mux.Vars(r)["int"]
	rule, err := loadRule(ruleId)
	if err != nil {
		return
	}
	renderTemplate(w, "rule.html", rule)
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/rule/{int}", ruleHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
