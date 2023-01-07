FROM golang:1.19-alpine as build-base

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/go-app .

# ===========

FROM alpine:3.16.2

COPY --from=build-base /app/out/go-app /app/go-app

ENV PORT=":2565"

ENV DATABASE_URL=postgres://drwpotzv:w2RbhmLcpNNQCTSwCNZLcIyhTKdbDrQ6@john.db.elephantsql.com/drwpotzv

CMD ["/app/go-app"]