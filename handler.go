package main

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
  "strconv"
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
  return
}

func (ed *edgeServer) getTeamHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab teamId from url
  teamName := ""
  if teamName = chi.URLParam(r, "name"); teamName == "" {
    http.Error(w, errors.New("no teamName in url.").Error(), http.StatusInternalServerError)
    log.Infof("Error in grabbing teamName from url. line 124. getTeamHandler(). \nbody: %v", r.Body)
    return
  }

  // call rpc method passing context and team from req
  team, err := ed.getTeam(r.Context(), teamName)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 132. getTeamHandler(). \nerr: %v", err.Error())
    return
  }

  successfulResponse := &TeamResponse{Message: "team retrieved", Success: true, Team: team}
  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 142. getTeamHandler(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(marshalledResp)
  return
}

func (ed *edgeServer) upsertProjectHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  user := r.Context().Value("user").(User)

  // grab teamId from url
  teamId := ""
  if teamId = chi.URLParam(r, "teamID"); teamId == "" {
    http.Error(w, errors.New("no teamID in url.").Error(), http.StatusInternalServerError)
    log.Infof("Error in grabbing teamID from url. line 157. upsertProjectHandler(). \nbody: %v", r.Body)
    return
  }

  // grab project from body
  // extract project from json body.
  var newProject ProjectRequest
  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in reading request body. line 167. createTeamHandler(). \nbody: %v", r.Body)
    return
  }
  newProject.TeamId = teamId
  // unmarshal json body into project request struct
  err = json.Unmarshal(reqBody, &newProject)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in unmarshalling body. line 175. upsertProjectHandler(). \nbody: %v", reqBody)
    return
  }

  // call rpc method
  err = ed.upsertProject(r.Context(), &newProject.Project, user.Id, teamId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 132. upsertProjectHandler(). \nerr: %v", err.Error())
    return
  }

  // send appropriate response and return
  successfulResponse := &Response{Message: "project upserted", Success: true}
  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 193. upsertProjectHandler(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(201)
  w.Write(marshalledResp)
  return
}

func (ed *edgeServer) getTeamsByUserHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab userId from url
  userId := ""
  if userId = chi.URLParam(r, "id"); userId == "" {
    http.Error(w, errors.New("no userId in url.").Error(), http.StatusInternalServerError)
    log.Infof("Error in grabbing userId from url. line 215. getTeamsByUserHandler(). \nbody: %v", r.Body)
    return
  }

  // call rpc method
  teams, err := ed.getTeamsByUser(r.Context(), userId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 223. getTeamsByUser(). \nerr: %v", err.Error())
    return
  }

  // send appropriate response and return
  successfulResponse := &TeamsResponse{Message: "teams retrieved", Success: true, Teams: teams}
  if len(teams) == 0 {
    successfulResponse.Message = "teams empty"
  }

  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 233. getTeamsByUser(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(marshalledResp)
  return
}

func (ed *edgeServer) getMyTeamsHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  // grab user from context that was passed from auth middleware, mostly just want id for rpc methods.
  user := r.Context().Value("user").(User)

  // call rpc method
  teams, err := ed.getMyTeams(r.Context(), user.Id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 253. getMyTeams(). \nerr: %v", err.Error())
    return
  }

  // send appropriate response and return
  successfulResponse := &TeamsResponse{Message: "teams retrieved", Success: true, Teams: teams}
  if len(teams) == 0 {
    successfulResponse.Message = "teams empty"
  }

  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 267. getMyTeams(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(marshalledResp)
  return

}

func (ed *edgeServer) getTeamsHandler(w http.ResponseWriter, r *http.Request) {
  log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

  page := int64(1)
  limit := int64(10)
  role := ""
  level := int64(0)
  technology := ""

  // extract query params
  // extract page to an int64
  if r.URL.Query()["page"] != nil {
    temp, err := strconv.Atoi(r.URL.Query()["page"][0])
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Infof("Error in grpc method. line 291. getTeams(). \nerr: %v", err.Error())
      return
    }
    page = int64(temp)
  }
  // extract limit to an int64
  if r.URL.Query()["limit"] != nil {
    temp, err := strconv.Atoi(r.URL.Query()["limit"][0])
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Infof("Error in grpc method. line 291. getTeams(). \nerr: %v", err.Error())
      return
    }
    limit = int64(temp)
  }
  // extract role to a string
  if r.URL.Query()["role"] != nil {
    role = r.URL.Query()["role"][0]
  }
  // extract level to an int64
  if r.URL.Query()["level"] != nil {
    temp, err := strconv.Atoi(r.URL.Query()["level"][0])
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Infof("Error in grpc method. line 317. getTeams(). \nerr: %v", err.Error())
      return
    }
    level = int64(temp)
  }
  // extract technology to a string
  if r.URL.Query()["technology"] != nil {
    technology = r.URL.Query()["technology"][0]
  }

  // call rpc method
  teams, err := ed.getTeams(r.Context(), role, technology, page, limit, level)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Infof("Error in grpc method. line 312. getTeams(). \nerr: %v", err.Error())
    return
  }

  // send appropriate response and return
  successfulResponse := &TeamsResponse{Message: "teams retrieved", Success: true, Teams: teams}
  if len(teams) == 0 {
    successfulResponse.Message = "teams empty"
  }

  marshalledResp, err := json.Marshal(successfulResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.WithField("error", err).Error("marshall error")
    log.Infof("Error in marshalling successful response. line 323. getTeams(). \nerr: %v", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(marshalledResp)
  return
}

// Routes to be implemented
/*
delete: "/v1/teams/{id}/members/{user_id}", delete member
delete: "/v1/teams/{id}", delete team
delete: "/v1/teams/{teamId}/members/{memberId}", remove member
get: "/{api}/teams" get list of teams


*/
