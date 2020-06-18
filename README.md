# A Simple Webapp

## Source Code

`fuck.go` simply fires up a webserver, and echoes back the parameter you pass on
the request's URL as a name to say hello to:

GET http://thisapp.as-deployed.somewhere/Joe

... will generate a simple "Hello, Joe!" page.

## Dockerfile

This just containerizes the app, in a 2 stage build process (so that the
resulting deployed container is basically *tiny*.

So now you can from this directory `docker run -it`, and then point your browser
at http://localhost:8999/Joe, and get your "Hello, Joe!" page that way.

## Kubernetes Manifest

It's not like I actually understand what I'm doing here.

1) Install a kubernetes cluster somewhere
2) `kubectl apply -f ./go-hello.yaml`

This will deploy *from my personal docker account* this same app to your
cluster. If you go investigate that account, you'll note that there are AMD64,
ARM64, and ARMv7 builds.
