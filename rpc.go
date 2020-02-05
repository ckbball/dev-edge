// this is where we actually call the gRPC services,
package main

import (
  teamSvc "github.com/ckbball/dev-team/pkg/api/v1"
  userSvc "github.com/ckbball/dev-user/pkg/api/v1"
)

func (ed *edgeServer) createTeam(ctx context.Context, team Team, creatorId string) error {
  team.Leader = creatorId
  _, err := teamSvc.NewTeamServiceClient(ed.teamSvcConn).CreateTeam(ctx, &teamSvc.TeamUpsertRequest{
    Api:  "v1",
    Team: team,
  })
  return err
}
