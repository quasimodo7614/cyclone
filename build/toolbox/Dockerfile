FROM alpine:3.8

LABEL maintainer="zhujian@caicloud.io"

WORKDIR /usr/bin/cyclone-toolbox

RUN apk add --no-cache coreutils

COPY ./build/toolbox/resolver-runner.sh ./bin/toolbox/fstream /usr/bin/cyclone-toolbox/
