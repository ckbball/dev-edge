package main

import (
  "net/http"
)

func (ed *edgeServer) registerUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (ed *edgeServer) createTeamHandler(w http.ResponseWriter, r *http.Request) {
  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  user := r.Context().Value("user").(User)

  // extract team from json body.

  // call rpc method
  r, err := ed.createTeam(r.Context(), team)
}
