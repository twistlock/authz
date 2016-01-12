FROM alpine:3.3
MAINTAINER Liron Levin <liron@twistlock.com>

VOLUME /var/lib/twistlock/policy.json
VOLUME /run/docker/plugins/

ADD ./authz_broker  /usr/bin/authz_broker

CMD ["/usr/bin/authz_broker"]