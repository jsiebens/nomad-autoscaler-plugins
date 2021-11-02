# nomad-autoscaler-plugins Dockerfile

#
# Build plugins in alpine container as upstream nomad-autoscaler does
#

FROM golang:alpine AS builder

ADD . /src
RUN apk add bash make \
    && cd /src \
    && make


#
# Install the plugins into the nomad-autoscaler images
#

FROM hashicorp/nomad-autoscaler:0.3.3

COPY --from=builder /src/bin/plugins/* /plugins/
