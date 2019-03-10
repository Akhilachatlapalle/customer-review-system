FROM golang:alpine
RUN apk add git
ENV buildpath=customer-review-system
ADD . /go/src/$buildpath
WORKDIR /go/src/$buildpath

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["go", "run", "main.go"]