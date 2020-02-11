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

var (
  apiVersion = "v1"
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
    Api:    apiVersion,
    Team:   svcTeam,
    UserId: creatorId,
  })
  if resp.Status == "error:duplicatename" {
    return errors.New("duplicate team name")
  } else if resp.Status == "error:maxteamcount" {
    return errors.New("max team count")
  }
  return err
}

func (ed *edgeServer) addMember(ctx context.Context, member *MemberRequest, ownerId string) error {
  // connect to service here
  conn, err := createConn(ctx, ed.teamSvcAddr)
  if err != nil {
    for err != nil {
      conn, err = createConn(ctx, ed.teamSvcAddr)
    }
  }

  resp, err := teamSvc.NewTeamServiceClient(conn).AddMember(ctx, &teamSvc.MemberUpsertRequest{
    Api:         apiVersion,
    TeamId:      member.TeamId,
    MemberId:    member.MemberId,
    MemberEmail: member.MemberEmail,
    Role:        member.Role,
    UserId:      ownerId,
  })
  if resp.Status == "error:exists" {
    return errors.New("member exists")
  } else if resp.Status == "error:maxmembercount" {
    return errors.New("max member count")
  }
  return err
}

func (ed *edgeServer) getTeam(ctx context.Context, teamName string) (*teamSvc.Team, error) {
  // connect to service here
  conn, err := createConn(ctx, ed.teamSvcAddr)
  if err != nil {
    for err != nil {
      conn, err = createConn(ctx, ed.teamSvcAddr)
    }
  }

  resp, err := teamSvc.NewTeamServiceClient(conn).GetTeamByTeamName(ctx, &teamSvc.GetByTeamNameRequest{
    Api:  apiVersion,
    Name: teamName,
  })
  if resp.Status == "error:missing" {
    return nil, errors.New("team doesn't exist")
  }

  return resp.Team, err

}

func (ed *edgeServer) upsertProject(ctx context.Context, project *Project, ownerId, teamId string) error {
  // connect to service here
  conn, err := createConn(ctx, ed.teamSvcAddr)
  if err != nil {
    for err != nil {
      conn, err = createConn(ctx, ed.teamSvcAddr)
    }
  }

  _, err = teamSvc.NewTeamServiceClient(conn).UpsertTeamProject(ctx, &teamSvc.ProjectUpsertRequest{
    Api:    apiVersion,
    TeamId: teamId,
    UserId: ownerId,
    Project: &teamSvc.Project{
      Description: project.Description,
      Languages:   project.Languages,
      Name:        project.Name,
      GithubLink:  project.GitLink,
      Complexity:  int32(project.Complexity),
      Duration:    int32(project.Duration),
    },
  })

  return err
}

// add a conversion function here
// func convertRestModelToRpc(interface, type string type of model to convert)

// func convertRpcModelToRest( interface, type string type of model to convert)
