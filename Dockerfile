FROM golang:alpine as builder
LABEL maintainer="Kushagra Indurkhya"
RUN apk update && apk add --no-cache git
ADD . /app
WORKDIR /app/server
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder /build ./build

ENV PORT=":8080"
RUN chmod +x ./main
EXPOSE 8080
CMD ./main