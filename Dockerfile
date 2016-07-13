FROM scratch
MAINTAINER Loggi "dev@loggi.com"
ADD main /
EXPOSE 8889
ENTRYPOINT ["/main"]
CMD ["-server", "-port", "8889"]
