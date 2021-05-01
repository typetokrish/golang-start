FROM golang:1.16-alpine
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/welcome-app
ARG GROUP_ID
ARG USER_ID
RUN addgroup -g $GROUP_ID app
RUN adduser -u $USER_ID -G app  -h /home/app -D $APP_USER 
RUN mkdir -p $APP_HOME
COPY ./ ${APP_HOME}
RUN chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER
WORKDIR $APP_HOME

EXPOSE 8010
RUN go mod download

RUN go mod tidy 
RUN go mod vendor
ENTRYPOINT go run main.go