# oauth-server

gin을 이용하여 인증서버를 구축합니다.  
redis와 의존됩니다.

# Stack

1. go:1.16.3

# 시퀀스 다이어그램

# History

## 프로젝트 생성 방법

1. `go mod init oauth-server`
1. `go get -u github.com/gin-gonic/gin`
1. `.vscode/launch.json` 에 `Launch Package` 템플릿으로 디버깅 환경 구축
