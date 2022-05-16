FROM golang:1.18-alpine
WORKDIR /app
# Copy mod and sum to workdir
COPY ../backend/go.mod ./
COPY ../backend/go.sum ./
RUN go mod download
# Copy all relevant .go files
COPY ../backend/*.go ./
RUN go build -o backend

CMD ["./backend"]

