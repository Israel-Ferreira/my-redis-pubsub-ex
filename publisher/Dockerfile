FROM golang:1.20-buster as build

WORKDIR /app

COPY . .

RUN go mod tidy


RUN CGO_ENABLED=0 GOOS=linux go build -o producer cmd/main.go


FROM scratch


WORKDIR /app


COPY --from=build /app/producer .

EXPOSE 4500

CMD ["./producer"]



