package GoGym

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	Delimiter        = "/"
	EscapedCharacter = "\\"
	End              = "$"
)

type Route struct {
	Uri      string
	Methods  []string
	Action   string
	Compiled Compiled
}

type Compiled struct {
	Tokens []Token
	RegExp string
}

type Token struct {
	Var     string
	Name    string
	Value   string
	IsParam bool
}

// ExtractTokens is a method extracts uri as tokens
func (r *Route) ExtractTokens(uri string) {
	expr := fmt.Sprintf("%s{(%sw+)%s}", EscapedCharacter, EscapedCharacter, EscapedCharacter)
	re, _ := regexp.Compile(expr)
	matches := re.FindAllStringSubmatch(uri, -1)
	var splitedStr []string
	if uri[0:1] == Delimiter {
		splitedStr = strings.Split(uri[1:], Delimiter)
	} else {
		splitedStr = strings.Split(uri, Delimiter)
	}
	tick := 0
	for _, str := range splitedStr {
		t := Token{}
		// If the length of matches equals 0, it means that there is
		// no variable left
		if len(matches) == 0 {
			t.Var = str
			t.Name = str
			t.Value = str
			t.IsParam = false
			r.Compiled.Tokens = append(r.Compiled.Tokens, t)
			continue
		}
		// Every time there is a variable match, the matched variable would
		// be pop out of slice, so it could be avoided to be compared again
		for _, match := range matches {
			tick++
			if str == match[0] {
				t.Var = match[0]
				t.Name = match[1]
				t.IsParam = true
				matches = append(matches[0:0], matches[1:]...)
				break
			} else {
				t.Var = str
				t.Name = str
				t.Value = str
				t.IsParam = false
			}
			//r.Compiled.Tokens = append(r.Compiled.Tokens, t)
		}
		r.Compiled.Tokens = append(r.Compiled.Tokens, t)
	}
}

// Compile is a method compiles tokens and regexp
func (r *Route) Compile(uri string) {
	// 除去字符串开头的"/"，为了好做正则
	var uriString string
	if uri[0:1] == Delimiter {
		uriString = uri[1:]
	} else {
		uriString = uri
	}
	var expression string
	splitString := strings.Split(uriString, Delimiter)
	for k, str := range splitString {
		token := r.Compiled.Tokens[k]
		if !token.IsParam {
			e := fmt.Sprintf("%s%s%s", EscapedCharacter, Delimiter, str)
			expression += e
		} else {
			e := fmt.Sprintf("%s%s(%sw+)", EscapedCharacter, Delimiter, EscapedCharacter)
			expression += e
		}
		if k == len(splitString)-1 {
			expression += End
		}
	}
	r.Compiled.RegExp = expression
}

// MatchAndAssign is a method that check if request uri matches defined uri
// and assign values to variables
func (r *Route) MatchAndAssign(uri string) {
	matched := r.Match(uri)
	if matched {
		r.AssignValuesToTokens(uri)
	}
}

// Match is a method to check if request uri meets regexp
func (r *Route) Match(uri string) bool {
	var matched bool
	regE, _ := regexp.Compile(r.Compiled.RegExp)
	matched = regE.MatchString(uri)
	splitString := strings.Split(uri, Delimiter)
	splitUri := strings.Split(r.Uri, Delimiter)
	// If request's path node is not equal as defined Uri, these two won't match
	if len(splitString) != len(splitUri) {
		matched = false
	}
	return matched
}

// AssignValuesToTokens is a method assigning values to tokens
func (r *Route) AssignValuesToTokens(uri string) {
	regE, _ := regexp.Compile(r.Compiled.RegExp)
	var stringMatches [][]string
	stringMatches = regE.FindAllStringSubmatch(uri, -1)
	// When the length of stringMatches[0] is greater than 1,
	// it means that there are variables in the path
	if len(stringMatches[0]) == 1 {
		return
	}
	// pointer equals 1 is because the first element of stringMatches is full request path,
	// not a variable
	pointer := 1
	for k, tkn := range r.Compiled.Tokens {
		if tkn.IsParam {
			r.Compiled.Tokens[k].Value = stringMatches[0][pointer]
			pointer++
		}
		// If pointer is greater than the length of stringMatches[0], there
		// would be no need to loop next round, because all params have been
		// assigned
		if pointer >= len(stringMatches[0]) {
			break
		}
	}
}
