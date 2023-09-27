FROM docker-reg.devops.xiaohongshu.com/shequ/centos:7

RUN yum install -y ca-certificates

ADD _output/bin/etcd-backup-operator /usr/local/bin/etcd-backup-operator
ADD _output/bin/etcd-restore-operator /usr/local/bin/etcd-restore-operator
ADD _output/bin/etcd-operator /usr/local/bin/etcd-operator

RUN adduser etcd-operator
USER etcd-operator
