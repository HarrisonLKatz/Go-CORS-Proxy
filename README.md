# Go CORS Proxy

Go CORS Proxy is exactly what it sounds like; it is an extremely simple CORS Proxy written in Go. Written in under 30 lines, this does not support advanced features such as further header manipulation or requests other than GET. Hosted on Repl.it at https://CORS-Proxy.codek1.repl.co

## Installation

Use [git](https://git-scm.com) to clone this repository.

```bash
git clone https://github.com/HarrisonLKatz/Go-CORS-Proxy.git
```

## Running

### Locally

```bash
go run CORSproxy.go
```
The port is :8000 by default for Repl.it compatibility, but this can be changed in the `CORSproxy.go` file.

### Repl.it

- First, import this Github repository (HarrisonLKatz/Go-CORS-Proxy) into a new Repl.
- Then, edit your .replit file to contain the following:
```
run = "go run CORSproxy.go"
language = "go"
```

## Usage

Simply prepend the server's URL to the URL you are trying to bypass CORS on. Do not forget to include protocal!

For example, if I was using https://CORS-Proxy.codek1.repl.co as the server and wanted to bypass CORS on example.com, I would link to https://CORS-Proxy.codek1.repl.co/https://example.com.

## Contributing
Pull requests are both welcomed and very much appreciated.

Please keep in mind that this is meant as a simple, lightweight, and easily readable proxy. Advanced features will not be considered for addition.

## License
[MIT](https://choosealicense.com/licenses/mit/)
