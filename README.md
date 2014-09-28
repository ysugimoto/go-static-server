# Go Static Server

Static server for development

## Usage

Clone or download this repostory and use `build/` binary for your operating system.

### Windows

```
./build/windows/go-static-server
```

### Linux

```
./build/linux/go-static-server
```

### Mac OS

```
./build/darwin/go-static-server
```

Build binary is supported 64-bit only.

## Supoorted extension(MimeType)

| Extension  | Description         | Serve                    |
| ---------- |:-------------------:|:------------------------:|
| .html/.htm | HTML Document       | text/html                |
| .css       | StyleSheet Document | text/css                 |
| .js        | JavaSCript Source   | application/javascript   |
| .json      | JSON Source         | application/json         |
| .gif       | GIF Image           | image/gif                |
| .jpg/.jpeg | JPEG Image          | image/jpeg               |
| .png       | PNG Image           | image/png                |
| .txt       | Plain Text          | text/plain               |
| (other)    | Octet-Stream        | application/octet-stream |
