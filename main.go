package main

import (
  "context"
  "fmt"
  "net/http"
  "os"
  "time"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/docgen"
  "github.com/go-chi/chi/middleware"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc"
)

const (
  port = "3000"
)

type edgeServer struct {
  teamSvcAddr string
  teamSvcConn *grpc.ClientConn
}

func main() {
  ctx := context.Background()
  log := logrus.New()
  log.Level = logrus.DebugLevel
  log.Formatter = &logrus.JSONFormatter{
    FieldMap: logrus.FieldMap{
      logrus.FieldKeyTime:  "timestamp",
      logrus.FieldKeyLevel: "severity",
      logrus.FieldKeyMsg:   "message",
    },
    TimestampFormat: time.RFC3339Nano,
  }
  log.Out = os.Stdout

  srvPort := port
  if os.Getenv("PORT") != "" {
    srvPort = os.Getenv("PORT")
  }
  addr := os.Getenv("LISTEN_ADDR")
  svc := new(edgeServer)
  mapEnv(&svc.teamSvcAddr, "TEAM_SERVICE_ADDR")

  connGRPC(ctx, &svc.teamSvcConn, svc.teamSvcAddr)

  r := chi.NewRouter()

  r.Use(middleware.Recoverer)
  r.Use(allowCors)

  r.Mount("/api/v1/teams", teamRouter())

  if *routes {
    fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
      ProjectPath: "github.com/ckbball/dev-edge",
      Intro:       "Welcome to the edge server of Dev Finder",
    }))
    return
  }

  var handler http.Handler = r
  handler = &logHandler{log: log, next: handler} // attach logger to handler

  log.Infof("Starting server at " + addr + ":" + srvPort)
  http.ListenAndServe(addr+":"+srvPort, handler)

}

func userRouter() chi.Router {}

func teamRouter(service *edgeServer) chi.Router {
  r := chi.NewRouter()
  r.Use(UserAuth)
  r.Post("", service.createTeamHandler)

  return r
}
