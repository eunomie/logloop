FROM scratch
MAINTAINER Yves Brissaud <eunomie@squarescale.com>
COPY app /app
ENTRYPOINT ["/app"]
