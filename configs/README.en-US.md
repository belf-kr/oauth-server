# Environmental variables

[한국어(KR)](./README.md) | [`English`](./README.en-US.md)

1. The variables that are all capital keys which are specified in `config` but the prefered priotiry is inserted `env` value.
1. To use kakao login, you must enter kakao application information at [kakao developers](https://developers.kakao.com/).
1. Be careful that token can be decrypted when JWT-related secret is leaked!

## Table legend

| Component     | Description                                                                                               |
| ------------- | --------------------------------------------------------------------------------------------------------- |
| Variable      | Environment Variable Name                                                                                 |
| dev           | Whether environmental variables are used in the development environment                                   |
| qa/prod       | Whether environmental variables are used in qa, production environments                                   |
| Default value | Value applied by default when no environment variable is determined using the system environment variable |
| Example       | List examples that can enter environment variable values                                                  |
| Explanation   | Description of environmental variables                                                                    |

## Table

| Variable              | dev | qa/prod | Default value | Example                 | Explanation                                                                                                     |
| --------------------- | :-: | :-----: | :-----------: | ----------------------- | --------------------------------------------------------------------------------------------------------------- |
| GO_ENV                | ✅  |   ✅    |      🤷‍♂️       | development, production | The value that sets `Go Execution Environment` and must have a value before starting the program starts.        |
| STAGES                | ✅  |   ✅    |      🤷‍♂️       | local, qa, prod         | To distinguish the deployment environment, the swagger document depends on the value.                           |
| SERVER_PORT           | ✅  |   ✅    |      🤷‍♂️       | 8080                    | Web Server HTTP listen port.                                                                                    |
| SWAGGER_HOSTNAME      | ✅  |   ✅    |      🤷‍♂️       | localhost, test.com     | Domain name used to troubleshoot CORS and request swagger documents.                                            |
| SWAGGER_PORT          | ✅  |   ✅    |      🤷‍♂️       | 8080, 443               | Port requesting swagger document.                                                                               |
| KAKAO_REST_API_KEY    | ✅  |   ✅    |      🤷‍♂️       |                         | Kakao login API required to use this is kakao API Key.                                                          |
| KAKAO_REDIRECT_URI    | ✅  |   ✅    |      🤷‍♂️       |                         | The address required to use the kakao login API and is called back upon completion of kakao login.              |
| JWT_ACCESS_SECRET     | ✅  |   ✅    |      🤷‍♂️       |                         | This is the key for encryption when issuing jwt access token.                                                   |
| JWT_REFRESH_SECRET    | ✅  |   ✅    |      🤷‍♂️       |                         | This is the key for encryption when issuing jwt refresh token.                                                  |
| AUTH_REDIRECT_URL     | ✅  |   ✅    |      🤷‍♂️       |                         | This is the address of the web app that will deliver the signed jwt token to queryString upon login completion. |
| MYSQL_MASTER_HOST     | ✅  |   ✅    |      🤷‍♂️       |                         | `DB Address` is the value used in the `MASTER environment`.                                                     |
| MYSQL_MASTER_PORT     | ✅  |   ✅    |      🤷‍♂️       |                         | `DB port` is the value used in the `MASTER environment`.                                                        |
| MYSQL_MASTER_USERNAME | ✅  |   ✅    |      🤷‍♂️       |                         | `DB Account Name` is the value used in the `MASTER environment`.                                                |
| MYSQL_MASTER_PASSWORD | ✅  |   ✅    |      🤷‍♂️       |                         | `Password for DB account` used in `MASTER environment`.                                                         |
| MYSQL_MASTER_DATABASE | ✅  |   ✅    |      🤷‍♂️       |                         | `DB name` is the value used in the `MASTER environment`.                                                        |
| REDIS_HOST            | ✅  |   ✅    |      🤷‍♂️       |                         | `DB Address`                                                                                                    |
| REDIS_PORT            | ✅  |   ✅    |      🤷‍♂️       |                         | `DB port`                                                                                                       |
| REDIS_PASSWORD        | ✅  |   ✅    |      🤷‍♂️       |                         | `Password for DB account`                                                                                       |
| REDIS_DATABASE        | ✅  |   ✅    |      🤷‍♂️       |                         | `DB name`                                                                                                       |
