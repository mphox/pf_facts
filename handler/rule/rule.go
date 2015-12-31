package rule

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mphox/pf_facts/handler"
	"github.com/mphox/pf_facts/model"
)

func RuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleId := mux.Vars(r)["int"]
	rule, err := model.LoadRule(ruleId)
	if err != nil {
		return
	}
	handler.RenderTemplate(w, "rule.html", rule)
	return
}
