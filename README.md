# Overview

OAuth 인증서버를 구축합니다.

1. JWT Token을 사용합니다.
1. 해당 앱을 사용하기 위해서는 kakao 개발자 센터에서 애플리케이션을 발급받아야 합니다. 발급 후 설정방법은 노션을 참고하도록 합니다.
1. API Spec는 `{url}/swagger/index.html` 으로 요청하여`swagger` 를 참고해주세요.
1. kakao login의 경우 콜백 때문에 테스트를 위해서는 웹앱이 필요합니다. [test/kakao-login](./test/kakao-login) 를 참고해주세요.

# Stack

1. go:1.16.3
1. vscode
1. gin
1. redis:6.2.4
1. docker

# Quick Start

## 개발 환경

1. vscode에서 `F5` 혹은 `go run main.go` 으로 시작합니다.
1. kakao developers에서 kakao application 정보 및 환경변수를 입력해주세요.

## 제품 시작

1. [config.prod.json](./config.prod.json) 에서 필요한 환경 변수를 설정합니다.
1. `docker-compose up -d` docker image 생성 후 컨테이너를 실행합니다.

> db 서비스가 초기화 후 웹 서버가 실행되어야하는데 `scratch` image라 [wait-for-it.sh](https://github.com/vishnubob/wait-for-it/) 를 실행하기가 쉽지 않았습니다. 근본적인 해결방법은 아니지만 local 환경에서 build 된 Dockerfile를 테스트하고 싶은 경우 재시작으로 해결하도록 합니다. local 이외 Prod는 k8s가 알아서 핸들링하기 때문에 상관이 없습니다.
