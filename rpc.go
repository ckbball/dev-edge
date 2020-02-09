// this is where we actually call the gRPC services,
package main

import (
  "context"
  "errors"
  "time"

  "google.golang.org/grpc"

  teamSvc "github.com/ckbball/dev-team/pkg/api/v1"
  //userSvc "github.com/ckbball/dev-user/pkg/api/v1"
)

func createConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
  conn, err := grpc.DialContext(ctx, addr,
    grpc.WithInsecure(),
    grpc.WithTimeout(time.Second*3))
  return conn, err
}

func (ed *edgeServer) createTeam(ctx context.Context, team *Team, creatorId string) error {
  // connect to service here
  conn, err := createConn(ctx, ed.teamSvcAddr)
  if err != nil {
    for err != nil {
      conn, err = createConn(ctx, ed.teamSvcAddr)
    }
  }

  team.Leader = creatorId
  protoProject := &teamSvc.Project{
    Description: team.Project.Description,
    Languages:   team.Project.Languages,
    Name:        team.Project.Name,
    GithubLink:  team.Project.GitLink,
    Complexity:  int32(team.Project.Complexity),
    Duration:    int32(team.Project.Duration),
  }
  svcTeam := &teamSvc.Team{
    Leader:     team.Leader,
    Name:       team.Name,
    OpenRoles:  int32(team.OpenRoles),
    Skills:     team.Skills,
    Size:       int32(team.Size),
    LastActive: int32(team.LastActive),
    Id:         team.Id,
    Project:    protoProject,
  }
  resp, err := teamSvc.NewTeamServiceClient(conn).CreateTeam(ctx, &teamSvc.TeamUpsertRequest{
    Api:  "v1",
    Team: svcTeam,
  })
  if resp.Status == "error:duplicatename" {
    return errors.New("duplicate team name")
  }
  return err
}

// add a conversion function here
// func convertRestModelToRpc(interface, type string type of model to convert)

// func convertRpcModelToRest( interface, type string type of model to convert)
