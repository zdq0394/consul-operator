FROM centos:7.4.1708
COPY make/release/consulops /opt/consul/operator

WORKDIR /opt/consul

ENTRYPOINT ["/opt/consul/operator", "consul"]
CMD ["--clusterdomain=cluster.local", "--concurrentworkers=3"]
