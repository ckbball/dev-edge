package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "time"

  // "github.com/pkg/errors"
  "github.com/sirupsen/logrus"
)

/*
func (ed *edgeServer) registerUserHandler(w http.ResponseWriter, r *http.Request) {

}
*/
func (ed *edgeServer) createTeamHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  timer2 := time.NewTimer(3 * time.Second)
  go func() {
    <-timer2.C
    log.Infof("Error: Request Timed Out")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    timeOutResponse := &Response{Message: "request timed out", Success: false}
    marshalledResp, err := json.Marshal(timeOutResponse)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.WithField("error", err).Error("marshall error")
      log.Infof("Error in marshalling timed out response. line 31. createTeamHandler(). \nerr: %v", err.Error())
      return
    }
    w.Write(marshalledResp)
    return
  }()

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
