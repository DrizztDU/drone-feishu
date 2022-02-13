# drone-feishu

[Drone](https://github.com/drone/drone) plugin for sending Feishu notifications.

> Inspired by https://github.com/appleboy/drone-telegram

## Usage

Execute from the working directory:

```shell
docker run --rm \
  -e PLUGIN_WEBHOOK=https://open.feishu.cn/open-apis/bot/v2/hook/... \
  -e DRONE_REPO=DrizztDU/drone-feishu \
  -e DRONE_REPO_LINK=https://github.com/DrizztDU/drone-feishu \
  -e DRONE_REPO_OWNER=DrizztDU \
  -e DRONE_REPO_NAME=drone-feishu \
  -e DRONE_COMMIT_SHA=e5e82b5eb3737205c25955dcc3dcacc839b7be52 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_MESSAGE='chore: update readme'  \
  -e DRONE_COMMIT_LINK=https://github.com/DrizztDU/drone-feishu/compare/master... \
  -e DRONE_COMMIT_AUTHOR=DrizztDU \
  -e DRONE_COMMIT_AUTHOR_EMAIL=xxxxxx@mail.com \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=https://cloud.drone.io/DrizztDU/drone-feishu \
  -e DRONE_TAG=1.0.0 \
  -e DRONE_BUILD_STARTED=1477550550 \
  -e DRONE_BUILD_FINISHED=1477550750 \
  ghcr.io/drizztdu/drone-feishu:linux-amd64
```
