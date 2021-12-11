# oauth-server

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

# 시퀀스 다이어그램

> 추가 예정

# 빠른 시작

1. vscode에서 `F5` 혹은 `go run main.go` 으로 시작합니다.

# 의존성

| 이름                                       | 명령어                                |
| ------------------------------------------ | ------------------------------------- |
| [gin](https://github.com/gin-gonic/gin)    | `go get -u github.com/gin-gonic/gin`  |
| [redis](https://github.com/go-redis/redis) | `go get github.com/go-redis/redis/v8` |
| [GORM](https://github.com/go-gorm/gorm)    | `go get -u gorm.io/gorm`              |
