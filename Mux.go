package GoGym

import (
	"fmt"
	"net/http"
)

type Mux struct {
	boss *Gym

	// mux *http.ServeMux
}

func (m *Mux) Prepare(g *Gym) {
	m.WhoIsYourBoss(g)
	// m.mux = http.NewServeMux()
}

func (m *Mux) WhoIsYourBoss(g *Gym) {
	m.boss = g
}

func (m *Mux) CallYourBoss() *Gym {
	return m.boss
}

func (m *Mux) GetMux() *Mux {
	return m
}

// ServeHTTP calls f(w, r).
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// do route logic
	if r.URL.Path == "/3" {
		fmt.Println("mux")
	}
	http.NotFound(w, r)
	return
}
