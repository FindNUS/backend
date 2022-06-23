FROM golang:1.18-alpine
WORKDIR /app
ARG PRODUCTION
ENV PRODUCTION=${PRODUCTION}
ARG MONGO_URI
ENV MONGO_URI=${MONGO_URI}
ARG RABBITMQ_URI
ENV RABBITMQ_URI=${RABBITMQ_URI}
ARG BONSAI_ES_URI
ENV BONSAI_ES_URI=${BONSAI_ES_URI}
# Copy mod and sum to workdir
COPY ./internal/item/go.mod ./
COPY ./internal/item/go.sum ./
RUN go mod download
# Copy all relevant .go files
COPY ./internal/item/*.go ./
RUN go build -o item

CMD ["./item"]

