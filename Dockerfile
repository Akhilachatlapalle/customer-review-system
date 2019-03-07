FROM golang:alpine
ENV buildpath=customer-review-system
ADD . /go/src/$buildpath
WORKDIR /go/src/$buildpath
CMD ["go", "run", "main.go"]