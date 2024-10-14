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
  project_name = "cairo"
  network_zone = "cairo.local"
  network_mask = "10.0.2.0"
  network_bits = 24
  network_range = "${local.network_mask}/${local.network_bits}"
  volume_source = "/qemu/qcow2/debian-sid-generic-amd64-daily-20241011-1897.qcow2"
  vm_props = {
    vm1 = {name = "${local.project_name}-vm1", ip = "10.0.2.11"}
    vm2 = {name = "${local.project_name}-vm2", ip = "10.0.2.12"}
  }
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

module "vms" {
  source                = "./vmdomain"
  for_each              = local.vm_props

  domain_name           = each.value.name
  domain_memory         = 1024
  domain_vcpu           = 1
  
  network_address       = each.value.ip
  network_bits          = local.network_bits
  network_gateway       = cidrhost(local.network_range, 1)
  network_name          = libvirt_network.network.name
  network_zone          = local.network_zone

  pool_name             = libvirt_pool.pool.name
  volume_source         = local.volume_source

  ssh_public_key_path   = file("~/.ssh/id_rsa.pub")
}

output "vm1_domain_name" {
  value = module.vms["vm1"].domain_name
}

output "vm2_network_address" {
  value = module.vms["vm2"].network_address
}

output "domain_names" {
  description = "List of all domain names"
  value = values(module.vms).*.domain_name
}

output "network_addresses" {
  description = "List of all network addresses"
  value = [for vm in module.vms : vm.network_address]
}

output "vm_details" {
  description = "Map of domain names to network addresses"
  value = {
    for key, vm in module.vms :
    vm.domain_name => vm.network_address
  }
}

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

# output "DomainName" {
#   value = module.vm1.domain_name
# }

# output "NetworkAddress" {
#   value = module.vm1.network_address
# }

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
