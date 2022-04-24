
FROM registry.access.redhat.com/ubi8/ubi:8.5
RUN yum -y update && yum clean all -y
RUN groupadd chatappgroup && \
    useradd -r -g chatappgroup chatappuser
USER chatappuser

LABEL maintainer="shariullas08@gmail.com"

COPY /build/chat-app /bin/
COPY ./migrations/ /tmp/migrations/
COPY ./asset/ /tmp/asset/