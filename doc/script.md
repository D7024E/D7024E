# Script
The start up script is "dock.sh", it uses the flags:


|Flag|Function
|----|----
|-u| Starts docker by firsrt initializing one node then starting the remaining nodes, as defined in the docker-compose file.
|-s| Rescales the docker network to a specified number of nodes, the scaling taken as a seperate user input prompt to the user.
|-d| Shuts down all active nodes.