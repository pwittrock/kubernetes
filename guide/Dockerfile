FROM phusion/baseimage

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -y
RUN apt-get -yy -q install nano
RUN apt-get -yy -q install dnsutils
RUN apt-get -yy -q install lxc

COPY service /etc/service/
COPY amd64/ /usr/local/bin
COPY cluster/ /cluster

CMD ["/sbin/my_init", "--", "bash", "-l"]
