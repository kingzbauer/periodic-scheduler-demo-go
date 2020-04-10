ARG BASE_IMG=golang:1.13.5-stretch
FROM $BASE_IMG

WORKDIR /home/app
COPY . .

RUN go mod tidy

ENTRYPOINT ["go"]
