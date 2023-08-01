# FROM golang:alpine AS builder

# LABEL stage=gobuilder

# ENV CGO_ENABLED 0
# ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct

# WORKDIR /app
# COPY ./bookstore .
# COPY ./etc .

# FROM alpine

# #RUN apk update --no-cache
# #RUN apk add --no-cache ca-certificates
# #RUN apk add --no-cache tzdata
# ENV TZ Asia/Taipei
# ENV LOCAL_HOST 192.168.31.132

# WORKDIR /app
# COPY --from=builder app/bookstore /app
# COPY --from=builder app/bookstore-api.yaml /app

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# CMD ["./bookstore", "-f", "bookstore-api.yaml"]


### Final Stage
# FROM golang:1.20-alpine
# ENV TZ Asia/Taipei
# WORKDIR /app
# RUN go install github.com/cosmtrek/air@latest
# COPY go.mod go.sum ./
# RUN go mod download
# CMD ["air", "-c", "./api/.air.toml"]

### Local 
FROM golang:1.20-alpine
RUN apk add tzdata
ENV TZ Asia/Taipei
WORKDIR /app
RUN apk del tzdata
RUN apk add curl
# 熱更新
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
CMD ["air", "-c", "./api/.air.toml"]