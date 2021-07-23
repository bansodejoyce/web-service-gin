FROM golang:latest AS build_base

# Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64

# Move to working directory /build
WORKDIR /tmp/web-service-gin

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
#RUN go build -o main .
RUN CGO_ENABLED=0 go build -o ./out/web-service-gin .

# Move to /dist directory as the place for resulting binary folder
#WORKDIR /dist

# Start fresh from a smaller image
FROM gcr.io/distroless/static

# Copy binary from build to main folder
# RUN cp /build/main .
COPY --from=build_base /tmp/web-service-gin/out/web-service-gin /

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
#CMD ["/dist/main"]

ENTRYPOINT [ "/web-service-gin" ]