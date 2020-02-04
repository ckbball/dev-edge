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
}

type Project struct {
  Description string
}
