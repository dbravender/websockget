websockget
==========

A command line websocket client.

Building:

    go get golang.org/x/net/websocket
    go build websockget.go

Example usage:

    websockget -headers="HEADER1: header1
    HEADER2: header2" -origin http://localhost/ ws://localhost:8082

Options:

     Usage of ./websockget:
     -headers="": A string of HTTP headers
     -origin="http://localhost/": Origin
