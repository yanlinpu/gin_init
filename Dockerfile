FROM golang:1.11.2

WORKDIR /app

COPY ./conf /app/conf
COPY ./bin/gin_init /app/gin_init

ENV GINENV=production
CMD ["./gin_init"]