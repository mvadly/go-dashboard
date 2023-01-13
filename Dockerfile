FROM golang:1.18

ENV TZ=Asia/Jakarta

WORKDIR /go-dashboard

COPY . /go-dashboard/

#RUN rm go.mod

#RUN rm go.sum

RUN ls -all

RUN apt-get update && apt-get install nano telnet traceroute -y

RUN go mod init go-dashboard

RUN go mod tidy

RUN go build -o /go-dashboard

EXPOSE 8080

CMD [ "/go-dashboard" ]