# netboot
Network boot software.

Getting started
===============
Run `git submodule init` to pull the files you'll need to get started.

iPXE boot images
================
To build the iPXE boot images, run `build_ipxe`. You will want to change the configuration files in the `config` directory.

Configuration files
-------------------
- `ca.crt` -- Certificate authority certificate. Default: [Jonathan Lung](http://www.heresjono.com)'s CA.
- `bootscript.ipxe` -- Certificate authority certificate. Default: Boot from <https://boot.heresjono.com:4747>.
- `ipxeclient.crt` -- Optional iPXE client certificate (recommended).
- `ipxeclient.key` -- Optional iPXE client key (recommended).
