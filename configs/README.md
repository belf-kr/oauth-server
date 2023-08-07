# 환경 변수

[`한국어(KR)`](./README.md) | [English](./README.en-US.md)

1. 모두 대문자 key인 환경변수의 경우 `config` 에도 지정되어 있고 `env` 로도 넣어준다면 `env` 값을 사용하게 됩니다.
1. kakao 로그인을 사용하기 위해서는 [kakao developers](https://developers.kakao.com/) 에서 kakao application 정보를 입력해야합니다.
1. JWT 관련 secret 유출 시 token을 복호화 할 수 있음으로 주의하도록 합니다!

## 표 범례

| 구성 요소     | 설명                                                                          |
| ------------- | ----------------------------------------------------------------------------- |
| Variable      | 환경 변수 이름                                                                |
| dev           | 환경 변수가 개발 환경에서 사용되는지 여부                                     |
| qa/prod       | 환경 변수가 qa, production 환경에서 사용되는지 여부                           |
| Default value | 시스템 환경 변수를 사용해 환경 변수를 정하지 않았을 때 기본적으로 적용되는 값 |
| Example       | 환경 변수 값으로 들어갈 수 있는 예시의 나열                                   |
| Explanation   | 환경 변수에 대한 설명                                                         |

## 표

| Variable              | dev | qa/prod | Default value | Example                 | Explanation                                                                           |
| --------------------- | :-: | :-----: | :-----------: | ----------------------- | ------------------------------------------------------------------------------------- |
| GO_ENV                | ✅  |   ✅    |      🤷‍♂️       | development, production | `Go 실행 환경` 을 설정하는 값이며 프로그램 시작 전 값이 있어야 합니다.                |
| STAGES                | ✅  |   ✅    |      🤷‍♂️       | local, qa, prod         | 배포 환경을 구분하기 위함이며, 해당 값에 따라서 swagger 문서가 달라집니다.            |
| SERVER_PORT           | ✅  |   ✅    |      🤷‍♂️       | 8080                    | 웹 서버 HTTP listen port 입니다.                                                      |
| SWAGGER_HOSTNAME      | ✅  |   ✅    |      🤷‍♂️       | localhost, test.com     | CORS 문제를 해결하기 위하여 사용되며 swagger 문서를 요청하는 도메인네임 입니다.       |
| SWAGGER_PORT          | ✅  |   ✅    |      🤷‍♂️       | 8080, 443               | swagger 문서를 요청하는 port 입니다.                                                  |
| KAKAO_REST_API_KEY    | ✅  |   ✅    |      🤷‍♂️       |                         | kakao 로그인 API를 사용하기 위해서 필요하며 kakao API Key 입니다.                     |
| KAKAO_REDIRECT_URI    | ✅  |   ✅    |      🤷‍♂️       |                         | kakao 로그인 API를 사용하기 위해서 필요하며 kakao 로그인 완료 시 콜백되는 주소입니다. |
| JWT_ACCESS_SECRET     | ✅  |   ✅    |      🤷‍♂️       |                         | jwt access token 발급 시 암호화 하기 위한 key 입니다.                                 |
| JWT_REFRESH_SECRET    | ✅  |   ✅    |      🤷‍♂️       |                         | jwt refresh token 발급 시 암호화 하기 위한 key 입니다.                                |
| AUTH_REDIRECT_URL     | ✅  |   ✅    |      🤷‍♂️       |                         | kakao 로그인 완료 시 서명된 jwt token을 queryString으로 전달 할 웹앱 주소입니다.      |
| MYSQL_MASTER_HOST     | ✅  |   ✅    |      🤷‍♂️       |                         | `DB 주소` 로 `MASTER 환경` 에서 사용되는 값입니다.                                    |
| MYSQL_MASTER_PORT     | ✅  |   ✅    |      🤷‍♂️       |                         | `DB port` 로 `MASTER 환경` 에서 사용되는 값입니다.                                    |
| MYSQL_MASTER_USERNAME | ✅  |   ✅    |      🤷‍♂️       |                         | `DB 계정명` 으로 `MASTER 환경` 에서 사용되는 값입니다.                                |
| MYSQL_MASTER_PASSWORD | ✅  |   ✅    |      🤷‍♂️       |                         | `DB 계정의 비밀번호` 로 `MASTER 환경` 에서 사용되는 값입니다.                         |
| MYSQL_MASTER_DATABASE | ✅  |   ✅    |      🤷‍♂️       |                         | `DB명` 으로 `MASTER 환경` 에서 사용되는 값입니다.                                     |
| REDIS_HOST            | ✅  |   ✅    |      🤷‍♂️       |                         | `DB 주소` 입니다.                                                                     |
| REDIS_PORT            | ✅  |   ✅    |      🤷‍♂️       |                         | `DB port` 입니다.                                                                     |
| REDIS_PASSWORD        | ✅  |   ✅    |      🤷‍♂️       |                         | `DB 계정의 비밀번호` 입니다.                                                          |
| REDIS_DATABASE        | ✅  |   ✅    |      🤷‍♂️       |                         | `DB명` 입니다.                                                                        |
