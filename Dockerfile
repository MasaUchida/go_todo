FROM golang:1.18

ENV TZ /user/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/todo
WORKDIR ${ROOT}

ENV GO111MODULE=on

COPY . .
EXPOSE 1323

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]