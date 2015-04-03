# Dockerfile for running the vapr bootserver with an ambassador from
# https://github.com/SojournLabs/ambassador.
FROM nginx
MAINTAINER jonathan lung <vapr@sojournlabs.com>
WORKDIR /vapr
ADD boot.conf /etc/nginx/sites-enabled/boot.conf
VOLUME ["/vapr/certstore", "/vapr/keys"]
CMD ["nginx", "-g", "daemon off;"]
