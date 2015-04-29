# servelement

I wrote servelement for fun. It serves the same purpose than [polyserve](https://github.com/PolymerLabs/polyserve) but it is written in Go.


`servelement` reads .bowerrc (if exists) to get the bower components directory. It reads bower.json to get my element name and serves 


## Installation

    go get -u github.com/rougepied/servelement

## Usage

Run `servelement`

```
cd my-element
servelement
```

## Browse files

Navigate to localhost:8080/my-element/demo.html

## Options

-p The TCP port to use for the web server