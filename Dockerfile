FROM alpine:3.7
MAINTAINER Liron Levin <liron@twistlock.com>

# Indicates basic authorization is enforced
ENV AUTHORIZER basic
# Indicates basic auditor type is used (log to console)
ENV AUDITOR basic
# Indicates audit logs are streamed to STDOUT
ENV AUDITOR-HOOK ""

VOLUME /var/lib/twistlock/policy.json
VOLUME /run/docker/plugins/

ADD ./bin/authz-broker  /usr/bin/authz-broker

CMD ["/usr/bin/authz-broker"]
