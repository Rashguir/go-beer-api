FROM golang:1.12.5-alpine3.9

MAINTAINER <rashguir@gmail.com> Nicolas Sicard

WORKDIR /go/src/app
RUN apk update; \
    apk add --no-cache -q -f \ 
        git


#COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

#CMD ["app"]

