FROM golang:1.10-alpine as builder
RUN apk add --no-cache git
RUN go get -v github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/ap8322/elm-golang-k8s-exercise/backend

# cache dependencies
COPY Gopkg.* ./
RUN dep ensure --vendor-only -v

COPY . .
RUN go install -v

FROM alpine:latest
COPY --from=builder /go/bin/backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]
