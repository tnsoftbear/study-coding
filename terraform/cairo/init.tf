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
