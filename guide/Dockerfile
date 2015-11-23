FROM phusion/baseimage


COPY service /etc/service/
COPY amd64/ /usr/local/bin

CMD ["/sbin/my_init", "--", "bash", "-l"]
