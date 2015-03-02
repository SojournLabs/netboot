FROM sojournlabs/gcc
MAINTAINER jonathan lung <vapr@sojournlabs.com>
RUN apt-get update && apt-get install -y \
    genisoimage \
    liblzma-dev \
    syslinux
RUN mkdir -p /vapr
WORKDIR /vapr
ENV PATH /vapr:$PATH
# Download ipxe
RUN git clone git://git.ipxe.org/ipxe.git
WORKDIR /vapr/ipxe/src
# Enable HTTPS and image signing.
RUN sed -i 's/undef\tDOWNLOAD_PROTO_HTTPS/define\tDOWNLOAD_PROTO_HTTPS/' config/general.h
RUN sed -i 's|//#define IMAGE_TRUST_CMD|#define\tIMAGE_TRUST_CMD|' config/general.h
# Do a preliminary build to save time later.
RUN make
WORKDIR /vapr
