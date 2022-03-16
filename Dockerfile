FROM node:17 as ui-builder
COPY . .
WORKDIR /ui
RUN npm ci
RUN npm run build

FROM golang:1.18-alpine3.15 as builder
WORKDIR /app/
COPY . .
RUN go mod tidy
RUN go build -o wrapit main.go


FROM alpine:3.15.0
ENV GITLAB_TOKEN=xxx
COPY --from=builder /app/wrapit /app/wrapit
COPY --from=ui-builder /ui/dist /app/dist
RUN adduser -D wrapit
RUN chown -R wrapit:wrapit /app
USER wrapit
CMD ["/app/wrapit","--static","/app/dist"]