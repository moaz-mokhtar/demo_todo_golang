FROM golang:latest

# set working directory
WORKDIR /app

# copy go mod and sum files
COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# copy the source code
COPY . .

# build the application
RUN go build -o main .

# expose port 8080 to the host
EXPOSE 8080

# command to run the executable
CMD ["./main"]
