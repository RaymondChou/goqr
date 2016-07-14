FROM scratch
MAINTAINER Loggi "dev@loggi.com"
ADD main /
ADD input/loggi_bit.png /input/loggi_bit.png
EXPOSE 8889
ENTRYPOINT ["/main"]
CMD ["-server", "-port", "8889"]
