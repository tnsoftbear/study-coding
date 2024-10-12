terraform {
    required_providers {
        libvirt = {
            source = "dmacvicar/libvirt"
            version = "0.8.0"
        }
    }
}

resource "libvirt_volume" "root" {
  name = "${var.domain_name}-root"
  pool = var.pool_name
  source = "https://cloud.debian.org/images/cloud/sid/daily/20241011-1897/debian-sid-generic-amd64-daily-20241011-1897.qcow2"
  format = "qcow2"
}

resource "libvirt_cloudinit_disk" "commoninit" {
  name = "${var.domain_name}-commoninit"
  pool = var.pool_name
  user_data = templatefile("${path.module}/files/cloud_init.cfg", {
    domain_name = var.domain_name
    network_zone = var.network_zone
    ssh_public_key = var.ssh_public_key_path
  })
  network_config = templatefile("${path.module}/files/network_config.cfg", {
    network_address = var.network_address
    network_bits = var.network_bits
    network_gateway = var.network_gateway
    network_zone = var.network_zone
  })
}

resource "libvirt_domain" "server" {
  name = var.domain_name
  memory = var.domain_memory
  vcpu = var.domain_vcpu
  cloudinit = libvirt_cloudinit_disk.commoninit.id
  
  network_interface {
    network_name = var.network_name
  }

  disk {
    volume_id = libvirt_volume.root.id
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