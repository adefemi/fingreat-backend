FROM golang:1.21 as builder

# Copy local code to the container image.
WORKDIR /fingreat_bk
COPY . .

# Build the command inside the container.
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./main.go

# Use a minimal alpine image for the final build to reduce image size
FROM alpine:3.14
COPY --from=builder /fingreat_bk/env.env /env.env
COPY --from=builder /fingreat_bk/server /server

# Run the service on container startup.
CMD ["/server"]