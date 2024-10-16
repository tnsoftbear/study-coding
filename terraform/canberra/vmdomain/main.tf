// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/volume
resource "libvirt_volume" "root" {
  name = "${var.domain_name}-root"
  pool = var.pool_name
  source = var.volume_source
  format = "qcow2"
}

// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/cloudinit
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

// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/domain
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
}
