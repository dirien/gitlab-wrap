FROM golang:1.17-alpine3.15 as builder
WORKDIR /app/
COPY . .
RUN go mod tidy
RUN go build -o wrapit main.go


FROM alpine:3.15.0
COPY --from=builder /app/wrapit /app/wrapit
COPY ./dist/ /app/dist/
RUN adduser -D wrapit
USER wrapit
CMD ["/app/wrapit"]