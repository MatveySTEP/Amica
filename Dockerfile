FROM golang:1.21-alpine as builder
LABEL stage=builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /project .

FROM alpine

COPY --from=builder /project /project

EXPOSE 8100

CMD [ "/project" ]