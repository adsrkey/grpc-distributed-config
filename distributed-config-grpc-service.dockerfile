FROM alpine:latest

RUN mkdir /app

COPY distributedConfigApp /app

CMD [ "/app/distributedConfigApp"]