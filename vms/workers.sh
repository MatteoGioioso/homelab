#!/usr/bin/env sh

worker_name=${1}

virt-install \
  --bridge=br0 \
  --name "${worker_name}" \
  --ram 4096 \
  --disk path=/var/lib/libvirt/images/"${worker_name}".qcow2,size=30 \
  --vcpus 1 \
  --cdrom /var/lib/libvirt/isos/ubuntu-20.04.2-live-server-amd64.iso
