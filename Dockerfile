FROM golang:1.21.5-bullseye

WORKDIR /src
COPY src .

RUN go build -o app main.go

EXPOSE 8080

ENTRYPOINT ["./app"]



# # Step 1: Use an official Go image as a parent image
# FROM golang:1.21.5-bullseye as builder

# # Step 2: Set the working directory inside the container
# WORKDIR /src

# # Step 3: Copy the local package files to the container's workspace
# COPY . .

# # Step 4: Download any necessary dependencies
# # This step uses the new Go modules (requires Go 1.11+)
# RUN go mod download

# # Step 5: Build the application for a production environment
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# # Step 6: Use a minimal alpine image for the final stage
# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Step 7: Copy the pre-built binary file from the previous stage
# COPY --from=builder /app/main .

# # Step 8: Expose port 8080 for the application
# EXPOSE 8080

# # Step 9: Run the binary program produced by `go install`
# CMD ["./main"]
