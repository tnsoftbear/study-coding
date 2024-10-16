terraform {
  required_providers {
    libvirt = {
      source = "dmacvicar/libvirt"
      version = "0.8.0"
    }
  }
}

// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs
provider "libvirt" {
  uri = "qemu:///system"
}

locals {
  project_name = "canberra"
  network_zone = "canberra.local"
  network_mask = "10.0.3.0"
  network_bits = 24
  network_range = "${local.network_mask}/${local.network_bits}"
  volume_source = "/qemu/qcow2/debian-sid-generic-amd64-daily-20241011-1897.qcow2"
}

// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/network
resource "libvirt_network" "network" {
  name = local.project_name
  autostart = true
  mode = "nat"
  domain = local.network_zone
  addresses = [local.network_range]

  dns {
    enabled = true
    local_only = true
  }
}

// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/pool
resource "libvirt_pool" "pool" {
  name = local.project_name
  type = "dir"
  path = "/qemu/disk/${local.project_name}"
}

module "masters" {
  source = "./vmdomain"
  count = 1

  domain_name       = "${local.project_name}-master${count.index}"
  domain_memory     = 2048
  domain_vcpu       = 1
  
  network_name      = libvirt_network.network.name
  network_zone      = local.network_zone
  network_address   = cidrhost(local.network_range, sum([10, count.index]))
  network_bits      = local.network_bits
  network_gateway   = cidrhost(local.network_range, 1)

  pool_name = libvirt_pool.pool.name
  volume_source = local.volume_source

  ssh_public_key_path = file("~/.ssh/id_rsa.pub")
}

module "nodes" {
  source = "./vmdomain"
  count = 2

  domain_name       = "${local.project_name}-node${count.index}"
  domain_memory     = 2048
  domain_vcpu       = 1
  
  network_name      = libvirt_network.network.name
  network_zone      = local.network_zone
  network_address   = cidrhost(local.network_range, sum([20, count.index]))
  network_bits      = local.network_bits
  network_gateway   = cidrhost(local.network_range, 1)

  pool_name = libvirt_pool.pool.name
  volume_source = local.volume_source

  ssh_public_key_path = file("~/.ssh/id_rsa.pub")
}

resource "local_file" "inventory_file" {
  content = templatefile("./templates/inventory.tpl", 
    {
      masters = module.masters
      nodes = module.nodes
    }
  )
  filename = "./inventory.ini"
}
