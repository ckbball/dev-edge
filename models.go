package main

import (
  teamSvc "github.com/ckbball/dev-team/pkg/api/v1"
)

type User struct {
  Id       string
  Username string
  Email    string
  Password string
}

type Team struct {
  Leader     string   `json:"leader"`
  Name       string   `json:"name"`
  OpenRoles  int      `json:"open_roles"`
  Skills     []string `json:"skills"`
  Size       int      `json:"size"`
  Id         string   `json:"id"`
  Project    Project  `json:"project"`
  LastActive int      `json:"last_active"`
  Members    []Member `json:"members"`
}

type Project struct {
  Description string   `json:"description"`
  Languages   []string `json:"languages"`
  Name        string   `json:"name"`
  GitLink     string   `json:"git_link"`
  Complexity  int      `json:"complexity"`
  Duration    int      `json:"duration"`
}

type Member struct {
  MemberId    string `json:"member_id"`
  MemberEmail string `json:"member_email"`
  Role        string `json:"role"`
}

type TeamRequest struct {
  Team Team   `json:"team"`
  Api  string `json:"api"`
  Id   string
}

type TeamResponse struct {
  Team    *teamSvc.Team `json:"team"`
  Message string        `json:"message"`
  Success bool          `json:"success"`
}

type TeamsResponse struct {
  Teams   []*teamSvc.Team `json:"teams"`
  Message string          `json:"message"`
  Success bool            `json:"success"`
}

type MemberRequest struct {
  TeamId      string `json:"team_id"`
  Api         string `json:"api"`
  MemberId    string `json:"member_id"`
  MemberEmail string `json:"member_email"`
  Role        string `json:"role"`
}

type ProjectRequest struct {
  Project Project
  Api     string
  TeamId  string
}

type FetchRequest struct {
  Team   Team
  Api    string
  UserId string
  Name   string
  // for GetTeams
  Page       int
  Limit      int
  Role       string
  Level      int
  Technology string
}

type Response struct {
  Message string
  Success bool
}
