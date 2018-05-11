FROM golang:alpine as build-env 

ADD . /src 

RUN cd /src && go build -o goapp


FROM alpine

COPY --from=build-env /src/goapp /traing-go

ENTRYPOINT [ "/traing-go" ]