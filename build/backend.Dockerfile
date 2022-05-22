FROM golang:1.18-alpine
WORKDIR /app
ARG PRODUCTION
ENV PRODUCTION=${PRODUCTION}
ARG FIREBASE_KEY
ENV FIREBASE_KEY=${FIREBASE_KEY}
# Copy mod and sum to workdir
COPY ./internal/backend/go.mod ./
COPY ./internal/backend/go.sum ./
RUN go mod download
# Copy all relevant .go files
COPY ./internal/backend/*.go ./
RUN go build -o backend

CMD ["./backend"]

