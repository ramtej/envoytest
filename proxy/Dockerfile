FROM lyft/envoy:latest
RUN apt-get update && apt-get -q install -y curl dnsutils
ADD ./service.pb /etc/service.pb
ADD ./hello-service.pb /etc/hello-service.pb
CMD /usr/local/bin/envoy -c /etc/envoy.json --service-cluster front-proxy
