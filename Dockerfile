FROM alpine:3.3
MAINTAINER Liron Levin <liron@twistlock.com>

VOLUME /var/lib/twistlock/policy.json
VOLUME /run/docker/plugins/

ADD ./twistlock_authz_plugin         /var/lib/twistlock/twistlock_authz_plugin

CMD ["/var/lib/twistlock/twistlock_authz_plugin"]