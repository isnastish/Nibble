FROM golang:1.23 AS build-env

WORKDIR /nibble/service/

ADD . /nibble/service/

# RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o /go/bin/service/txn github.com/isnastish/nibble/service/

# FROM golang:1.23-alpine3.21 AS run-env

# COPY --from=build-env /go/bin/service/txn /txn/

# EXPOSE 3030 

CMD [ "ls", "-all" ]