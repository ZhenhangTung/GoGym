package GoGym

import (
	"regexp"
)

type Route struct {
	uri      string
	methods  []string
	action   interface{}
	compiled [string]interface{}
}

func (r *Route) urlMatches(url *url.URL, route *Route) {
	expression := r.getCompiled("regexp")
	regexp.Match(expression, url)

}

func (r *Route) getCompiled(key string) (interface{}, bool) {
	v, ok := r.compiled[key]
	return v, ok
}
