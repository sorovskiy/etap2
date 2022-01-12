package handlers

import (
	"encoding/json"
	"etap2/domain"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RepoInt interface {
	GetState(application string) (domain.APIResp, error)
	SetState(resp domain.APIResp) error
	UpdateState(options domain.APIResp, resp domain.APIResp) error
}

type SetParams struct {
	Repo RepoInt
}

func NewParamsSetter(repo RepoInt) *SetParams {
	return &SetParams{
		Repo: repo,
	}
}

func (p *SetParams) Routes() chi.Router {
	root := chi.NewRouter()
	root.Use(middleware.Logger)

	root.Get("/api/getstate", p.GetState)
	root.Post("/api/savestate", p.SetState)

	return root
}

func (p *SetParams) GetState(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var options domain.APIResp
	err = json.Unmarshal(body, &options)
	if err != nil {
		log.Println("Unmarshall error")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Json unmarshall error")
		return
	}

	resp, err := p.Repo.GetState(options.Application)
	if err != nil {
		log.Printf("Client \"%s\" is absent in db", options.Application)
		_, _ = io.WriteString(w, "No such client in db")
		return
	}

	fmt.Printf("Client's data: %+v\n", resp)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	_ = enc.Encode(resp)
}

func (p *SetParams) SetState(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var options domain.APIResp
	err = json.Unmarshal(body, &options)
	if err != nil {
		log.Println("Unmarshall error")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Json unmarshall error")
		return
	}

	resp, err := p.Repo.GetState(options.Application)
	if err != nil {
		options.Version = 1
		err = p.Repo.SetState(options)
		if err != nil {
			log.Println("Can't write to db")
		}
		log.Println("Created the record in db with new client")
		_, _ = io.WriteString(w, "Created the record in db with new client")
		return
	}

	if resp.Param1 == options.Param1 && resp.Param2 == options.Param2 {
		log.Printf("Version of \"%s\" hadn't been updated: data is same", options.Application)
		_, _ = io.WriteString(w, "Version of client hadn't been updated")
		return
	}

	err = p.Repo.UpdateState(options, resp)
	_, _ = io.WriteString(w, "Version of client had been updated")
	log.Printf("Version of \"%s\" had been updated", options.Application)
}
