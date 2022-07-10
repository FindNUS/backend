FROM golang:1.18-alpine
WORKDIR /app
ARG PRODUCTION
ENV PRODUCTION=${PRODUCTION}
ARG FIREBASE_KEY
ENV FIREBASE_KEY=${FIREBASE_KEY}
ARG FIREBASE_KEY_ID
ENV FIREBASE_KEY_ID=${FIREBASE_KEY_ID}
ARG MONGO_URI
ENV MONGO_URI=${MONGO_URI}
ARG RABBITMQ_URI
ENV RABBITMQ_URI=${RABBITMQ_URI}
ARG BONSAI_ES_URI
ENV BONSAI_ES_URI=${BONSAI_ES_URI}
ARG MONGO_URI
ENV MONGO_URI=${MONGO_URI}
# Copy mod and sum to workdir
COPY ./internal/lookout/go.mod ./
COPY ./internal/lookout/go.sum ./
RUN go mod download
# Copy all relevant .go files
COPY ./internal/lookout/*.go ./
RUN go build -o lookout

CMD ["./lookout"]

