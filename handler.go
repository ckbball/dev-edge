package main

import (
  "encoding/json"
  "ioutil"
  "net/http"
)

/*
func (ed *edgeServer) registerUserHandler(w http.ResponseWriter, r *http.Request) {

}
*/
func (ed *edgeServer) createTeamHandler(w http.ResponseWriter, r *http.Request) {
  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  user := r.Context().Value("user").(User)

  // extract team from json body.
  var newTeam TeamRequest
  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {

  }

  // unmarshal json body into team request struct
  err = json.Unmarshal(reqBody, &newTeam)
  if err != nil {

  }

  // call rpc method passing context and team from req
  result, err := ed.createTeam(r.Context(), newTeam.Team)

  w.WriteHeader(http.StatusCreated)
}
