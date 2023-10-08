#!/usr/bin/env fish

sudo virt-install --name fedora-app2 \
    --memory 3074 --cpu host --vcpus 3 --graphics none \
    --os-type linux --os-variant fedora38  --import \
    --disk /var/lib/libvirt/images/fedora-app2.qcow2,format=qcow2,bus=virtio \
    --network bridge=virbr1,model=virtio \
    --cloud-init
