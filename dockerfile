FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main ./cmd

# ------------------------------------------
FROM scratch

COPY --from=builder /app/main /cmd/main

EXPOSE 3000
CMD [ "/cmd/main" ]
# ----------------------------------------
#   command to build: docker build -t abdulrahimom/infuzio:latest . 
#   command to build: docker build -t infuzio:latest . 
#  command to run container: 