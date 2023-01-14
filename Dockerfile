FROM golang:1.18

ENV TZ=Asia/Jakarta

WORKDIR /api-go-dashboard

COPY . /api-go-dashboard/

#RUN rm go.mod

#RUN rm go.sum

RUN ls -all

RUN apt-get update && apt-get install nano telnet traceroute -y

RUN go mod init api-go-dashboard

RUN go mod tidy

RUN go build -o /api-go-dashboard

EXPOSE 8080

CMD [ "/api-go-dashboard" ]