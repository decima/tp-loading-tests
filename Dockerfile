FROM golang:1.20-alpine as backend
RUN mkdir /app
COPY go.mod /app
COPY go.sum /app
WORKDIR /app

RUN go mod download
COPY . /app

RUN CGO_ENABLED=0 go build -o chargeApp .


FROM alpine as runner
RUN mkdir -p /app
ENV GIN_MODE=release
WORKDIR /app
COPY --from=backend /app/chargeApp ./chargeApp
RUN chmod a+x chargeApp
ENTRYPOINT /app/chargeApp