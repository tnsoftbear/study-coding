variable "domain_name" {
  type = string
  description = "The domain name assigned to the virtual machine."
}

variable "domain_memory" {
  type = string
  description = "Amount of memory (in MB) allocated to the virtual machine."
}

variable "domain_vcpu" {
  type = number
  description = "Number of virtual CPUs allocated to the virtual machine."
}

variable "pool_name" {
  type = string
  description = "Name of the storage pool where the virtual machine disk will be stored."
}

variable "network_name" {
  type = string
  description = "Name of the virtual network to which the virtual machine will be connected."
}

variable "network_zone" {
  type = string
  description = "DNS zone or domain for the virtual network, used for internal domain name resolution."
}

variable "network_address" {
  type = string
  description = "Base IP address of the virtual network."
}

variable "network_bits" {
  type = number
  description = "The number of bits for the network mask, defining the network size."
}

variable "network_gateway" {
  type = string
  description = "The IP address of the network gateway for external traffic routing."
}

variable "ssh_public_key_path" {
  type = string
  description = "Path to the SSH public key file, used for SSH access to the virtual machine."
}
