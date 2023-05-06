# syntax=docker/dockerfile:1

FROM golang:1.20 as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY api ./api/
COPY backend ./backend/
COPY main.go .
COPY go.mod .
COPY go.sum .
COPY temp/openapi.yaml .
RUN mkdir /app/temp
RUN go get .

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /compute

FROM scratch AS runtime
COPY --from=build /compute ./
COPY --from=build app/openapi.yaml ./api/openapi.yaml

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/compute"]
