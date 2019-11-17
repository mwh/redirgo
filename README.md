A simple web server that serves redirects from flat files

The server can be started with

    redirgo /path/to/files 12345

to read its data from /path/to/files and bind on localhost:12345. Any
incoming request will be mapped to the corresponding file inside that
path, which should contain an absolute HTTP(S) URL. The server will
generate a 302 redirect to that address, or a 404 if no such file
exists.

The intended use is as a minimal backing server proxied behind something
like nginx or lighttpd to serve redirections from a sub-area of a web
site, backed by a flat-file database and without any changes to the web
server configuration to modify the mapping of redirections dynamically.

Build with

    go build
