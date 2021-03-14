###################################
# STEP 1 build executable binary
###################################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/amazon-clone-account-service
COPY . .

# Fetch dependencies using go get.
RUN go get -d -v 

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/amazon-clone-account-service

###################################
# STEP 2 build a small image
###################################
FROM scratch

# Copy static executable.
COPY --from=builder /go/bin/amazon-clone-account-service /go/bin/amazon-clone-account-service

# Run the binary.
ENTRYPOINT [ "/go/bin/amazon-clone-account-service" ]