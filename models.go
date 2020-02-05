package main

type User struct {
  Id       string
  Username string
  Email    string
  Password string
}

type Team struct {
  Leader    string
  Name      string
  OpenRoles int
  Skills    []string
  Size      int
  Id        string
  Project   Project
  Members   []Member
}

type Project struct {
  Description string
  Languages   []string
  Name        string
  GitLink     string
  Complexity  int
  Duration    int
}

type TeamRequest struct {
  Team Team
  Api  string
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
