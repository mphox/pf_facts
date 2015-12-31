package root

import (
	"math/rand"
	"net/http"
	"strconv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	ruleId := strconv.Itoa(rand.Intn(2) + 1)
	http.Redirect(w, r, "/rule/"+ruleId, http.StatusFound)
}
