FROM golang:bullseye

# create build directory for source code
WORKDIR /build

# copies everything to build directories
COPY . .

# install go deps
RUN go mod download

# Build
RUN go build -o /authentication-service

EXPOSE 9990

CMD [ "/authentication-service" ]
