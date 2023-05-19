FROM golang:latest AS build

WORKDIR /usr/local/showstart
COPY . /usr/local/showstart
RUN mkdir bin
RUN make build


FROM ubuntu:20.04
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive \
    GOGC=800
WORKDIR /usr/local
RUN mkdir /usr/local/backend
RUN mkdir /usr/local/backend/conf
RUN mkdir /usr/local/backend/bin
COPY --from=build /usr/local/showstart/bin /usr/local/backend/bin/
COPY --from=build /usr/local/showstart/conf/config.yaml /usr/local/backend/conf
WORKDIR /usr/local/backend/bin

EXPOSE 8000
CMD ["/usr/local/backend/bin/showstart","-c","../conf/config.yaml"]