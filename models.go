package main

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
}

type Project struct {
  Description string   `json:"description"`
  Languages   []string `json:"languages"`
  Name        string   `json:"name"`
  GitLink     string   `json:"git_link"`
  Complexity  int      `json:"complexity"`
  Duration    int      `json:"duration"`
}

type TeamRequest struct {
  Team Team   `json:"team"`
  Api  string `json:"api"`
  Id   string
}

type MemberRequest struct {
  TeamId      string
  Api         string
  MemberId    string
  MemberEmail string
  Role        string
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
