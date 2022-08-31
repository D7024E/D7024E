# Docker
How to setup a docker, can be found here; https://docs.docker.com/engine/reference/commandline/stack_deploy/

## Docker Swarm

### Setup: 
> docker swarm init
### Deploy: 
> docker stack deploy --compose-file docker-compose.yml vossibility
### Change number of replications
> docker service scale vossibility_kademliaNodes=<new number of replicas>
### Exit: 
> docker swarm leave --force

## Docker-Compose

### Setup & Deploy 
> docker compose up -d
### Status
> docker compose ps -a
### Exit 
> docker compose down

