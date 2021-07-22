#!/usr/bin/env sh

virt-install --bridge=br0 \
  --name pfsense \
  --ram 2048 \
  --disk path=/var/lib/libvirt/images/pfsense.qcow2,size=10 \
  --vcpus 1 \
  --cdrom /var/lib/libvirt/isos/pfSense-CE-2.5.2-RELEASE-amd64.iso