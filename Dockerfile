FROM golang:1.21-alpine AS backend-builder

WORKDIR /build
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend/ .
RUN go build -o main .

FROM scratch

COPY --from=backend-builder /build/main .
ENTRYPOINT ["/main"]
