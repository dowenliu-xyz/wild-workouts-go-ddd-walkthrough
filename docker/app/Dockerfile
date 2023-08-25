ARG GOPROXY

FROM golang:1.21.0
ARG GOPROXY

ENV GOPROXY=$GOPROXY

RUN go install github.com/cespare/reflex@latest

COPY reflex.conf /

COPY start.sh /

RUN chmod +x /start.sh

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
