#A Domain-Specific Language (DSL) is a computer language that's targeted to a particular kind of problem
#rather than a general purpose language that's aimed at any kind of software problem.
#baseImage 
FROM golang:1.22.3

WORKDIR /myAscii

#Docker Optimization...
#Read about why / slash https://docs.docker.com/reference/dockerfile/#copy
COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -v -o main main.go

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
#Just for Documentation / we still need to 
EXPOSE 8080



CMD [ "./main" ] 
#or just --> CMD ./main
