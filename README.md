# netboot
Network boot software.

Getting started
===============
Run `git submodule init` to pull the files you'll need to get started.

iPXE boot images
================
To build the iPXE boot images, run `build_ipxe`. You will want to change the configuration files in the `config` directory.

iPXE boot server
================
To build the iPXE boot server, run `build_ipxe_server`. You will want to change the configuration files in the `ipxe_server.config` directory. To run the iPXE server, run `run_ipxe_server`. To use the proxy, use `run_ipxe_server_proxy`.


Configuration files
-------------------
- certstore/`ca.crt` -- Certificate authority certificate. Default: [Jonathan Lung](http://www.heresjono.com)'s CA.
- `ipxe_client.config/bootscript.ipxe` -- Certificate authority certificate. Default: Boot from <https://boot.heresjono.com:4747>.
- `ipxe_client.config/ipxeclient.key` -- Optional iPXE client key with corresponding certificate in certstore (recommended).
- `ipxe_server_proxy.config/ipxeserver.key` -- iPXE server key with corresponding certificate in certstore.