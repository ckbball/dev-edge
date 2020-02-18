# dev-edge

The edge http REST API that exposes the appropriate services to the outside world in addition to validating auth requests.

- It currently exposes routes that enable access to the backend microservices.
- It is written in Go and runs on Docker.

In this README:

- [Routes](#routes)
- [Repositories](#repositories)
- [Future](#future)

## Routes

The current routes available are:

- POST: /api/v1/teams ~ This endpoint takes a team object and creates a new Team.
- POST: /api/v1/teams/{id}/members ~ This endpoint takes member information and adds a new member to the corresponding Team.
- POST: /api/v1/teams/{id}/projects ~ This endpoint takes a project object and upserts it to the corresponding Team.
- GET: /api/v1/teams/users/{id} ~ This endpoint returns all the teams a user owns or is a member of.
- GET: /api/v1/teams/{name} ~ This endpoint returns the full information(team info, project info, member info) of the team
specified by team name.
- GET: /api/v1/teams ~ This endpoint returns a list of teams that match the query params

## Repositories

This edge service currently only relies upon the Team service here: https://github.com/ckbball/dev-team

## Future

Future changes and additions to this service include:

- Endpoints to expose the User service.
- Logging and other necessary metrics.
- Kubernetes deployment files.
