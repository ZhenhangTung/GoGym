package GoGym

import (
	"fmt"
	"net/url"
	// "regexp"
)

type Route struct {
	uri      string
	methods  []string
	action   interface{}
	compiled Compiled
}

func (r *Route) urlMatches(u *url.URL, route *Route) {
	expression := r.getCompiledRegPattern()
	fmt.Println(u)
	fmt.Println(expression)
	// regexp.MatchString(expression, )

}

func (r *Route) getCompiledRegPattern() string {
	return r.compiled.pattern
}

type Compiled struct {
	pattern string
}
