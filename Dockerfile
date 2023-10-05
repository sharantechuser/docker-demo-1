FROM golang:latest
RUN mkdir /app
ADD . /app
# EXPOSE 4000
WORKDIR /app
RUN go build -o docker1 .
CMD [ "/app/docker1" ]

ENV USER_NAME Test