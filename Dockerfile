FROM golang:1.23.0-alpine AS builder
WORKDIR /usr/local/src
RUN apk add --no-cache bash git gcc musl-dev
COPY ["go.mod", "go.sum", "./"] 
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . ./

RUN ["templ", "generate"]

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app .

FROM alpine AS runner
COPY --from=builder /usr/local/src/bin/app ./
COPY .env ./
COPY ./static ./static
CMD ["/app"]
