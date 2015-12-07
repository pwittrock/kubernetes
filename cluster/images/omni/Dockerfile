FROM ubuntu:trusty

RUN apt-get update

RUN apt-get -y install lxc
RUN apt-get -y install curl

RUN curl -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.1.2/bin/linux/amd64/kubectl
RUN chmod 755 /usr/local/bin/kubectl

COPY cluster.sh /usr/local/bin

ENTRYPOINT ["cluster.sh"]

CMD ["up"]
