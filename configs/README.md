# 환경 변수

1. 모두 대문자 key인 환경변수의 경우 `config` 에도 지정되어 있고 `env` 로도 넣어준다면 `env` 값을 사용하게 됩니다.
1. kakao 로그인을 사용하기 위해서는 [kakao developers](https://developers.kakao.com/) 에서 kakao application 정보를 입력해야합니다.

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

| Variable              | dev | qa/prod | Default value | Example                 | Explanation                                                            |
| --------------------- | :-: | :-----: | :-----------: | ----------------------- | ---------------------------------------------------------------------- |
| GO_ENV                | ✅  |   ✅    |      🤷‍♂️       | development, production | `Go 실행 환경` 을 설정하는 값이며 프로그램 시작 전 값이 있어야 합니다. |
| STAGES                | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| SERVER_PORT           | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| SWAGGER_HOSTNAME      | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| SWAGGER_PORT          | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| KAKAO_REST_API_KEY    | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| KAKAO_REDIRECT_URI    | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| JWT_ACCESS_SECRET     | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| JWT_REFRESH_SECRET    | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| AUTH_REDIRECT_URL     | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| MYSQL_MASTER_HOST     | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| MYSQL_MASTER_PORT     | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| MYSQL_MASTER_USERNAME | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| MYSQL_MASTER_PASSWORD | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| MYSQL_MASTER_DATABASE | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| REDIS_HOST            | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| REDIS_PORT            | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| REDIS_PASSWORD        | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
| REDIS_DATABASE        | 🚫  |   ✅    |       ?       | ?                       | ?                                                                      |
