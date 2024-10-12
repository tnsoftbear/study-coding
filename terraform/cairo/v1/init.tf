terraform {
    required_providers {
        libvirt = {
            source = "dmacvicar/libvirt"
            version = "0.8.0"
        }
    }
}

provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_network" "network" {
  name = "cairo"
  autostart = true
  mode = "nat"
  domain = "cairo.local"
  addresses = ["10.0.2.0/24"]

  dns {
    enabled = true
    local_only = true
  }
}

resource "libvirt_pool" "pool" {
  name = "cairo"
  type = "dir"
  path = "/qemu/disk/cairo"
}

resource "libvirt_volume" "server_root" {
  name = "server_root"
  pool = libvirt_pool.pool.name
  source = "https://cloud.debian.org/images/cloud/sid/daily/20241011-1897/debian-sid-generic-amd64-daily-20241011-1897.qcow2"
  format = "qcow2"
}

data "template_file" "user_data" {
  template = file("${path.module}/files/cloud_init.cfg")
  vars = {
    domain_name = "cairo.local"
  }
}

data "template_file" "network_config" {
  template = file("${path.module}/files/network_config.cfg")
}

resource "libvirt_cloudinit_disk" "commoninit" {
  name = "server_cloudinit"
  pool = libvirt_pool.pool.name
  user_data = data.template_file.user_data.rendered
  network_config = data.template_file.network_config.rendered
}

resource "libvirt_domain" "server" {
  name = "cairo-vm1"
  memory = "1024"
  vcpu = 1
  cloudinit = libvirt_cloudinit_disk.commoninit.id
  
  network_interface {
    network_name = libvirt_network.network.name
  }

  disk {
    volume_id = libvirt_volume.server_root.id
  }

  console {
    type = "pty"
    target_port = "0"
    target_type = "serial"
  }

  console {
    type = "pty"
    target_type = "virtio"
    target_port = "1"
  }

  graphics {
    type = "spice"
    listen_type = "address"
    autoport = true
  }
}