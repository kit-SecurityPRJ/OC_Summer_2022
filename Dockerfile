FROM ubuntu

COPY hosts coredns Corefile /

RUN apt-get update && apt-get upgrade -y

CMD ["./coredns"]

EXPOSE 53 53/udp
