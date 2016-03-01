FROM ubuntu:15.10
#COPY cluster-manager /usr/bin/
RUN apt-get update && \
    apt-get install --no-install-recommends -y redis-server zookeeper spiped
RUN dd if=/dev/urandom bs=32 count=1 of=keyfile
RUN apt-get install -y iproute2
ENV ZK_ID 1
ENV CLUSTER_SIZE 3
COPY zk.sh /usr/bin/
CMD ["cluster-manager"]
