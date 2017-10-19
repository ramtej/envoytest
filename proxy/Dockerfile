FROM lyft/envoy:latest
RUN apt-get update && apt-get -q install -y curl dnsutils
RUN touch /var/log/envoy.log
CMD /usr/local/bin/envoy -l 1 -c /etc/envoy.json --service-cluster front-proxy
