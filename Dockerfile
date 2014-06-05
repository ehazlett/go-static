from stackbrew/debian:jessie
maintainer evan hazlett <ejhazlett@gmail.com>
add go-static /usr/local/bin/go-static
run echo "hello from go-static" > /var/tmp/index.html
expose 3000
entrypoint ["/usr/local/bin/go-static"]
cmd ["-static-dir=/var/tmp"]
