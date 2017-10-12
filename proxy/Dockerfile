FROM lyft/envoy:latest
RUN apt-get update && apt-get -q install -y curl dnsutils
CMD /usr/local/bin/envoy -c /etc/envoy.json --service-cluster front-proxy
