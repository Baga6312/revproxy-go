
# RevProxy_GO

## clone of TinyPR 

Just a simple reverse proxy writting in golang  

The Purpose of this is to bypass warnings when tunneling with ngrok or localtunnel 

## Installation 
installing requried packages first

`go mod tidy `

## Usage 

run dummy containers for initializing traffic  

`make run-containers`

running the proxy 

`make run-proxy-server`

stop the containers 

`make stop`

help for usage 

`make help`

