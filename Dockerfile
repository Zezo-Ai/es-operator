FROM registry.opensource.zalan.do/library/alpine-3.12:latest
MAINTAINER Team Lens @ Zalando SE <team-lens@zalando.de>

# add binary
ADD build/linux/es-operator /

ENTRYPOINT ["/es-operator"]
