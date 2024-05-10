FROM golang:1.22-alpine
WORKDIR /app
# ARG PRODUCTION
# ENV PRODUCTION=${PRODUCTION}
# ARG FIREBASE_KEY
# ENV FIREBASE_KEY=${FIREBASE_KEY}
# ARG FIREBASE_KEY_ID
# ENV FIREBASE_KEY_ID=${FIREBASE_KEY_ID}
# ARG MONGO_URI
# ENV MONGO_URI=${MONGO_URI}
# ARG RABBITMQ_URI
# ENV RABBITMQ_URI=${RABBITMQ_URI}
# Copy mod and sum to workdir
COPY ./internal/backend/go.mod ./
COPY ./internal/backend/go.sum ./
RUN go mod download
# Copy all relevant .go files
COPY ./internal/backend/*.go ./
RUN go build -o backend

CMD ["./backend"]

