# Overview

1. JWT Token을 기반으로 구축 된 인증서버 입니다.
1. API Spec는 `{url}/swagger/index.html` 의 `swagger` 를 참고해주세요.
1. kakao login의 경우 콜백 때문에 테스트를 위해서는 웹앱이 필요합니다. [test/kakao-login](./test/kakao-login) 를 참고해주세요.
1. 환경 변수 범례의 경우 [configs](./configs) 를 참고해주세요.
1. 이외 컨테이너 실행과 같은 명령어는 vscode `tasks.json` 에도 정의되어 있음으로 편하게 task로 실행하세요! 😎

# Stack

1. go:1.16.3
1. vscode
1. gin
1. redis:6.2.4
1. mysql:5.7.16
1. docker

# Quick Start

## 개발 환경

1. [config.dev.json](./configs/config.dev.json) 에서 필요한 환경 변수를 설정합니다.
1. `docker-compose up -d` 으로 db server와 같은 개발에 필요한 컨테이너를 올립니다.
1. vscode의 `실행 및 디버그` 에서 `Server` 으로 실행 혹은 터미널에 `export GO_ENV=development && go run main.go` 으로 시작합니다.

## 제품 시작

1. [config.prod.json](./configs/config.prod.json) 에서 필요한 환경 변수를 설정합니다.
1. `docker-compose up -d` docker image 생성 후 컨테이너를 실행합니다.

> db 서비스가 초기화 후 웹 서버가 실행되어야하는데 `scratch` image라 [wait-for-it.sh](https://github.com/vishnubob/wait-for-it/) 를 실행하기가 쉽지 않았습니다. 근본적인 해결방법은 아니지만 local 환경에서 build 된 Dockerfile를 테스트하고 싶은 경우 재시작으로 해결하도록 합니다. local 이외 Prod는 k8s가 알아서 핸들링하기 때문에 상관이 없습니다.

# Other

## swagger 확인 및 업데이트 방법

swagger가 탑재되어 있음으로 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) 에서 swagger 문서를 확인하실 수 있습니다.  
이외, swagger 업데이트 명령어가 vscode `tasks.json` 에 정의되어 있음으로 편하게 task로 실행하세요! 😎

> 직접 명령어로 수행 하시려면 아래의 명령어를 사용하세요.
>
> ```shell
> export PATH=$(go env GOPATH)/bin:$PATH
> swag i
> ```

## live reloading로 실행 방법

javascript, python, ruby와 같은 인터프리터 언어로 작업하는 데 익숙한 사람들은 특히 golang에서 live reloading을 갈망합니다.  
이를 위해 `tasks.json` 에 live reloading를 위한 명령어를 정의해 놓았습니다. 해당 기능을 사용하기 위해서 아래의 패키지를 다운받아주세요.

```shell
go get github.com/codegangsta/gin
```

> live reloading로 속도감 있게 개발하는 것은 좋지만 아쉬운 점은 breakpoint을 사용하기 위해서는 필요할 때마다 process를 attach 해줘야합니다.  
> 이유는 reloading 마다 빌드되어 실행되는 파일이 달라서 process를 계속 추적할 수 없습니다. (아직 debug tool에서 지원되지 않는 것으로 확인됨)  
> 만약, live reloading에서 breakpoint가 필요하면 `launch.json` 에 정의된 `Attach to Gin` 를 실행하세요! (이럴 줄 알고 미리 정의해 놓았습니다)  
> reloading시 연결이 끈어지지만 빠른 개발을 위해 live reloading 진행하다가 필요할 떄 debug를 해나가는 전략이 더 유리할 것으로 기대합니다.
