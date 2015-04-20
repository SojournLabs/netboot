#!/bin/bash
if [ -d /btrfs ]; then
    echo "File system already initialized?"
    exit 1
fi
mkdir /btrfs
cryptsetup -v luksFormat /dev/sda --key-size=512;
cryptsetup luksOpen /dev/sda encroot
mkfs.btrfs -L ROOT /dev/mapper/encroot
mount /dev/mapper/encroot /btrfs
btrfs subvolume create /btrfs/var
btrfs subvolume create /btrfs/etc
btrfs subvolume create /btrfs/home
btrfs subvolume create /btrfs/root
cp -R /var/* /btrfs/var
cp -R /etc/* /btrfs/etc
cp -R /home/* /btrfs/home
cp -R /root/* /btrfs

mkdir /btrfs/home/docker
mkdir /btrfs/home/etcd
chown docker:docker -R /btrfs/home/docker
chown etcd:etcd -R /btrfs/home/etcd


# Set up nameservers
cp /usr/share/resolv.conf.head /btrfs/etc