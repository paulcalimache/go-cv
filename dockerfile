FROM golang:1.22-alpine as build
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o gocv

FROM chromedp/headless-shell:latest
RUN apt-get update; apt install dumb-init -y
# ENTRYPOINT ["dumb-init", "--"]
COPY --from=build /usr/src/app /app
WORKDIR /app
ENTRYPOINT ["dumb-init", "--", "./gocv"]