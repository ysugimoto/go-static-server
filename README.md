# Go Static Server

Static server for development

## Usage

There is two way to use it:

- by cloning this project
- by using the `go get` command

### Cloning this project

Clone or download this repostory and use `build/` binary for your operating system.

#### Windows

```
./build/windows/go-static-server.exe
```

#### Linux

```
./build/linux/go-static-server
```

#### Mac OS

```
./build/darwin/go-static-server
```

Build binary is supported 64-bit only.

### Using the `go get` command

Simply run the command

```bash
 go get github.com/ysugimoto/go-static-server
```
in your shell.

If you have added the `$GOPATH/bin` to your path, you are able to run the command `go-static-server`.  
Otherwise, you'll find the compiled version in the `$GOPATH/bin/go-static-server` directory.

## Command Line options

- `-p` Listen port number ( default is 8888)
- `-d` Change DocumentRoot (default is current directory)

for example:

```
go-static-server -p 9999 -d ~/server
```

Will listen `0.0.0.0:9999` and document root is `$HOME/server`. 

## Supported extension (MimeTypes)

| Extension  | Description         | Serve mimetypes          |
| ---------- |:-------------------:|:------------------------:|
| .html/.htm | HTML Document       | text/html                |
| .css       | StyleSheet Document | text/css                 |
| .js        | JavaScript Source   | application/javascript   |
| .json      | JSON Source         | application/json         |
| .gif       | GIF Image           | image/gif                |
| .jpg/.jpeg | JPEG Image          | image/jpeg               |
| .png       | PNG Image           | image/png                |
| .txt       | Plain Text          | text/plain               |
| (other)    | Octet-Stream        | application/octet-stream |

