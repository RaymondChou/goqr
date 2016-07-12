goqr
====

A fast qrcode generate write with google golang.

Build:
    go build main.go
    
And use (will generate images in output directory):
    goqr -data=sometext,anothertext,moretext

Or start a http server:
    goqr -server

And make a request:
    http://localhost:8889?data=myqrcodecontent


Forked from https://github.com/RaymondChou/goqr
