# Overview

[한국어(KR)](./README.md) | [`English`](./README.en-US.md)

1. It is an authentication server built based on JWT Token.
1. For API Spec, refer to `swagger` in `{url}/swagger/index.html`.
1. In the case of kakao login, we need a web app for testing because of callback. Please refer to [test/kakao-login](./test/kakao-login).
1. For environment variable legends, see [configs](./configs).
1. Other commands such as executing containers are also defined in vscode `tasks.json`, so please run them as tasks comfortably! 😎

# Stack

1. go:1.16.3
1. vscode
1. gin
1. redis:6.2.4
1. mysql:5.7.16
1. docker

# Quick Start

## Development environment

1. [config.dev.json](./configs/config.dev.json) sets the required environment variables.
1. Raise the container needed for development such as db server with `docker-compose up -d`.
1. Start with `Export GO_ENV=development && go run main.go` in `Run and Debug` in vscode as `Server` or in the terminal.

## Launch Product

1. Set the required environment variables in [config.prod.json](./configs/config.prod.json).
1. Run the container after creating the `docker-compose up -d` docker image.

> The web server should be running after the db service is initialized, but it was not easy to run [wait-for-it.sh](https://github.com/vishnubob/wait-for-it/) because of the `scratch` image. If you want to test a built Dockerfile in a local environment, though not a fundamental solution, try restarting it. Prod other than local doesn't matter because k8s handles it on its own.

# Other

## How to check and update swaggers

With swagger installed, you can check the swagger document at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).  
In addition, the swagger update command is defined in vscode `tasks.json`, so please run it as a task! 😎

> Use the command below to run directly.
>
> ```shell
> export PATH=$(go env GOPATH)/bin:$PATH
> swag i
> ```

## How to run with live reloading

People who are used to working in interpreter languages like javascript, python, ruby, especially crave live reloading in golang.  
To this end, we have defined a command for live reloading in `tasks.json`. Please download the package below to use this function.

```shell
go get github.com/codegangsta/gin
```

> It is good to develop it quickly with live reloading, but unfortunately, in order to use breakpoint, the process must be attached whenever necessary.  
> The reason is that each reloading builds and executes different files, so the process cannot be tracked continuously. (Not yet supported by the debug tool)
> If live reloading requires a breakpoint, run `Attach to Gin` as defined in `launch.json`! (I knew this would happen, so I defined it in advance.)  
> Although the connection is disconnected during reloading, we expect that the strategy to proceed with live reloading for rapid development and debug when necessary will be more advantageous.
