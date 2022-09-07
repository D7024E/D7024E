# Docker
Decision to use just docker swarm was made since it was recommended and stil utilizes the docker-compose. 

## Docker Build 
> docker build . -t kadlab

## Docker list ip addresses
> docker ps -q | xargs -n 1 docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} {{ .Name }}' | sed 's/ \// /'

## Docker-Compose
Docker compose information can be found here, https://docs.docker.com/compose/reference/.
### Setup & Deploy 
> docker compose up -d
### Status
> docker compose ps -a
### Exit 
> docker compose down

## Docker Swarm
How to setup a docker swarm, can be found here, https://docs.docker.com/engine/reference/commandline/stack_deploy/.
### Setup: 
> docker swarm init
### Deploy: 
> docker stack deploy --compose-file docker-compose.yml vossibility
### Change number of replications
> docker service scale vossibility_kademliaNodes=<new number of replicas>
### Exit: 
> docker swarm leave --force

