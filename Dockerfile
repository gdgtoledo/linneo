############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/gdgtoledo/linneo

COPY . .

# Build the binary.
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -o /go/bin/linneo
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/linneo /go/bin/linneo
# Run the linneo binary.
ENTRYPOINT ["/go/bin/linneo"]