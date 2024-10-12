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

locals {
  project_name = "cairo"
  network_zone = "cairo.local"
}

resource "libvirt_network" "network" {
  name = local.project_name
  autostart = true
  mode = "nat"
  domain = local.network_zone
  addresses = ["10.0.2.0/24"]

  dns {
    enabled = true
    local_only = true
  }
}

resource "libvirt_pool" "pool" {
  name = local.project_name
  type = "dir"
  path = "/qemu/disk/${local.project_name}"
}

locals {
  vm_names = {
    vm1 = {name = "${local.project_name}-vm1", ip = "10.0.2.11"}
    vm2 = {name = "${local.project_name}-vm2", ip = "10.0.2.12"}
  }
}

module "vms" {
  source = "./vmdomain"

  for_each = local.vm_names
  domain_name     = each.value.name
  domain_memory = "1024"
  domain_vcpu = 1
  
  network_name = libvirt_network.network.name
  network_zone = local.network_zone
  network_address = each.value.ip
  network_bits = 24
  network_gateway = "10.0.2.1"

  pool_name = libvirt_pool.pool.name

  ssh_public_key_path = file("~/.ssh/id_rsa.pub")
}

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

# module "vms" {
#   count = 2
#   source = "./vmdomain"

#   domain_name = "${local.project_name}-vm${count.index + 1}"
#   domain_memory = "1024"
#   domain_vcpu = 1
  
#   network_name = libvirt_network.network.name
#   network_zone = local.network_zone
#   network_address = cidrhost("10.0.2.0/24", sum([11, count.index]))
#   network_bits = 24
#   network_gateway = "10.0.2.1"

#   pool_name = libvirt_pool.pool.name

#   ssh_public_key_path = file("~/.ssh/id_rsa.pub")
# }

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

# module "vm1" {
#   source = "./vmdomain"

#   domain_name = "${local.project_name}-vm1"
#   domain_memory = "1024"
#   domain_vcpu = 1
  
#   network_name = libvirt_network.network.name
#   network_zone = local.network_zone
#   network_address = "10.0.2.11"
#   network_bits = 24
#   network_gateway = "10.0.2.1"

#   pool_name = libvirt_pool.pool.name

#   ssh_public_key_path = file("~/.ssh/id_rsa.pub")
# }

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
