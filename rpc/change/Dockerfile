# FROM golang:alpine AS builder

# LABEL stage=gobuilder

# ENV CGO_ENABLED 0
# ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct

# WORKDIR /app
# COPY ./change .
# COPY ./etc .

# FROM alpine

# #RUN apk update --no-cache
# #RUN apk add --no-cache ca-certificates
# #RUN apk add --no-cache tzdata
# ENV TZ Asia/Taipei

# WORKDIR /app
# COPY --from=builder app/change /app
# COPY --from=builder app/change.yaml /app

# CMD ["./change", "-f", "change.yaml"]

## Local
FROM golang:1.20-alpine
RUN apk add tzdata
ENV TZ Asia/Taipei
WORKDIR /app
RUN apk del tzdata
RUN apk add curl
# 熱更新
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
CMD ["air", "-c", "./api/.air.toml"]