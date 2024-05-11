# Build stage
FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

# Final stage
FROM build

RUN chmod +x ./app.exe

EXPOSE 8000 5432

CMD ["./app.exe"]