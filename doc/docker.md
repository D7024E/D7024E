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
> docker-compose down
### Add Nodes
> docker-compose scale kademliaNodes=<Total Number Of Nodes>
### Add Nodes And Update
> docker-compose up --scale kademliaNodes=<Total Number Of Nodes>


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

# CLI
In order to access a running nodes active terminal use the command:
> docker  attach <"the nodes id">

Beware that closing this terminal process will close the node. For testing purposes you can use this randomly genated kademlia hash.
> 6b5106626a8bcfa1d12a940294d286aa2ae0f54c

