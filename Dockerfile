FROM scratch
MAINTAINER Loggi "dev@loggi.com"
ADD main /
ENTRYPOINT ["/main"]
CMD ["-server", "-port", "8889"]
