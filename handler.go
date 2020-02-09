package main

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
  //"time"

  // "github.com/pkg/errors"
  "github.com/go-chi/chi"
  "github.com/sirupsen/logrus"
)

/*
func (ed *edgeServer) registerUserHandler(w http.ResponseWriter, r *http.Request) {

}
*/
func (ed *edgeServer) createTeamHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  user := r.Context().Value("user").(User)

  // extract team from json body.
  var newTeam TeamRequest
  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in reading request body. line 27. createTeamHandler(). \nbody: %v", r.Body)
    return
  }

  // unmarshal json body into team request struct
  err = json.Unmarshal(reqBody, &newTeam)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in unmarshalling body. line 35. createTeamHandler(). \nbody: %v", reqBody)
    return
  }

  // call rpc method passing context and team from req
  err = ed.createTeam(r.Context(), &newTeam.Team, user.Id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 43. createTeamHandler(). \nerr: %v", err.Error())
    return
  }

  successfulResponse := &Response{Message: "team created", Success: true}
  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 52. createTeamHandler(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  w.Write(marshalledResp)
}

func (ed *edgeServer) addMemberHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  user := r.Context().Value("user").(User)

  // grab teamId from url
  teamID := ""
  if teamID = chi.URLParam(r, "teamID"); teamID == "" {
    http.Error(w, errors.New("no teamID in url.").Error(), http.StatusInternalServerError)
    log.Infof("Error in grabbing teamID from url. line 71. addMemberHandler(). \nbody: %v", r.Body)
    return
  }

  // extract team from json body.
  var newMember MemberRequest
  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in reading request body. line 27. createTeamHandler(). \nbody: %v", r.Body)
    return
  }
  newMember.TeamId = teamID
  // unmarshal json body into team request struct
  err = json.Unmarshal(reqBody, &newMember)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in unmarshalling body. line 81. addMemberHandler(). \nbody: %v", reqBody)
    return
  }

  // call rpc method passing context and team from req
  err = ed.addMember(r.Context(), &newMember, user.Id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 89. addMemberHandler(). \nerr: %v", err.Error())
    return
  }

  successfulResponse := &Response{Message: "member added", Success: true}
  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 98. addMemberHandler(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  w.Write(marshalledResp)

}

func (ed *edgeServer) getTeamHandler(w http.ResponseWriter, r *http.Request) {
  //log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  //user := r.Context().Value("user").(User)
  return
}

// Routes to be implemented
/*
post: "/v1/teams/{id}/members",
delete: "/v1/teams/{id}/members/{member_number}",
post: "/v1/teams/{id}/project",
get: "/{api}/teams/users/{id}"
get: "/{api}/myteams"
get: "/{api}/teams"
get: "/{api}/teams/{name}"
*/
