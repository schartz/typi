# Typi
Typi is a dead simple collaborative text/code editor.

## Setting up development environment

 Install [Go](https://golang.org/doc/install)

 Clone repository somewhere in your $GOPATH

    git clone https://github.com/schartz/typi.git
    cd wherever/you/cloned/this

Now install node modules:

    yarn install
  
## Running and developing

    yarn run watch

 This uses ```nodemon``` to watch and build as you edit front end assets.
    
     yarn run build
     yarn run production
     
 These will build the javascript for development and production respectively

    go run *.go
    
 This will build and run the back end server
 
    go build -o typi *.go
    
 This will build a static binary named ```typi``` which can be run standalone.

## TODO
1) Come up with a logo.
2) Migrate to webpack.
3) Replace CSS workflow with SASS workflow.
4) Update mgo or ditch Mongo altogether, maybe.

## Roadmap
1) User accounts to store the documents.
2) Add 1 to 1 video calling

## Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other method.