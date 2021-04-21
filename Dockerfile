FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/boiler-plate-app
FROM alpine:latest AS production
COPY --from=builder /app/boiler-plate-app /app/boiler-plate-app
COPY --from=builder /app/application.env /app/application.env
# COPY --from=builder /app/application.env /app/application.env
EXPOSE 8080
CMD [ "/app/boiler-plate-app" ]

