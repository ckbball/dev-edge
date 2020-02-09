package main

import (
  "context"
  "flag"
  "fmt"
  "net/http"
  "os"
  "time"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/docgen"
  "github.com/pkg/errors"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc"
)

const (
  port = "3000"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

type edgeServer struct {
  teamSvcAddr string
  teamSvcConn *grpc.ClientConn
}

func main() {
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

  r := chi.NewRouter()

  r.Use(middleware.Recoverer)
  r.Use(allowCORS)

  r.Mount("/api/v1", teamRouter(svc))

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

  log.Fatal(http.ListenAndServe(addr+":"+srvPort, handler))
  return
}

func userRouter() chi.Router {
  r := chi.NewRouter()
  return r
}

func teamRouter(service *edgeServer) chi.Router {
  r := chi.NewRouter()
  r.Use(UserAuth)
  r.Post("/teams", service.createTeamHandler)
  // /teams/{id}

  r.Route("/teams/{teamID}", func(r chi.Router) {
    r.Post("/members", service.addMemberHandler)
  })
  r.Get("/{name}", service.getTeamHandler)

  return r
}

func mapEnv(target *string, envVar string) {
  v := os.Getenv(envVar)
  if v == "" {
    panic(fmt.Sprintf("environment variable %q not set", envVar))
  }
  *target = v
}

func connGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
  var err error
  *conn, err = grpc.DialContext(ctx, addr,
    grpc.WithInsecure(),
    grpc.WithTimeout(time.Second*3))
  if err != nil {
    panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
  }
}
