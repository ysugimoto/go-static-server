# Go Static Server

Static server for development

## Usage

Clone or download this repostory and use `build/` binary for your operating system.

Running:

```
./go-static-server
```

Listens at `0.0.0.0:8888` with document root being the same as binary directory itself.
See `Command Line options` for customization below.

### Windows

```
./build/windows/go-static-server.exe
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

## Command Line options

- `-b` Bind to interface (default is 0.0.0.0)
- `-p` Listen port number (default is 8888)
- `-d` Change DocumentRoot (default is current directory)

for example:

```
go-static-server -b 127.0.0.1 -p 9999 -d ~/server
```

Will listen `127.0.0.1:9999` and document root is `$HOME/server`.

