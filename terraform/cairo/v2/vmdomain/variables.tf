variable "domain_name" {
  type = string
  description = "The domain name assigned to the virtual machine."

  validation {
    condition     = length(var.domain_name) >= 3 && length(var.domain_name) <= 15
    error_message = "Domain name must be between 3 and 15 characters."
  }
}

variable "domain_memory" {
  type = number
  description = "Amount of memory (in MB) allocated to the virtual machine."
  default = 1024

  validation {
    condition     = var.domain_memory >= 512 && var.domain_memory <= 8192
    error_message = "Memory size must be between 512 MB and 8192 MB."
  }
}

variable "domain_vcpu" {
  type = number
  description = "Number of virtual CPUs allocated to the virtual machine."
  default = 1

  validation {
    condition     = var.domain_vcpu >= 1 && var.domain_vcpu <= 8
    error_message = "vCPU count must be between 1 and 8 units."
  }
}

variable "pool_name" {
  type = string
  description = "Name of the storage pool where the virtual machine disk will be stored."

  validation {
    condition     = length(var.pool_name) > 0 && can(regex("^[a-zA-Z0-9_-]+$", var.pool_name))
    error_message = "The pool name must not be empty and can only contain alphanumeric characters, hyphens, or underscores."
  }
}

variable "network_name" {
  type = string
  description = "Name of the virtual network to which the virtual machine will be connected."

  validation {
    condition     = length(var.network_name) > 0 && can(regex("^[a-zA-Z0-9_-]+$", var.network_name))
    error_message = "The network name must not be empty and can only contain alphanumeric characters, hyphens, or underscores."
  }
}

variable "network_zone" {
  type = string
  description = "DNS zone or domain for the virtual network, used for internal domain name resolution."

  validation {
    condition     = length(var.network_zone) > 0 && can(regex("^[a-zA-Z0-9.-]+$", var.network_zone))
    error_message = "The network zone must not be empty and can only contain alphanumeric characters, dots, or hyphens."
  }
}

variable "network_address" {
  type = string
  description = "Base IP address of the virtual network."

  validation {
    condition     = can(regex("^([0-9]{1,3}\\.){3}[0-9]{1,3}$", var.network_address))
    error_message = "The network address must be a valid IPv4 address (e.g., 192.168.1.10)."
  }
}

variable "network_bits" {
  type = number
  description = "The number of bits for the network mask, defining the network size."

  validation {
    condition     = var.network_bits >= 0 && var.network_bits <= 32
    error_message = "Network bits must be between 0 and 32."
  }
}

variable "network_gateway" {
  type = string
  description = "The IP address of the network gateway for external traffic routing."

  validation {
    condition     = can(regex("^([0-9]{1,3}\\.){3}[0-9]{1,3}$", var.network_gateway))
    error_message = "The network gateway must be a valid IPv4 address (e.g., 192.168.1.1)."
  }
}

variable "ssh_public_key_path" {
  type = string
  description = "Path to the SSH public key file, used for SSH access to the virtual machine."
  sensitive = true

  validation {
    condition     = can(regex("^ssh-(rsa|ed25519|ecdsa) [A-Za-z0-9+/=]+[ ]?.*", var.ssh_public_key_path))
    error_message = "The SSH public key must be in a valid format (ssh-rsa, ssh-ed25519, or ecdsa)."
  }
}

variable "volume_source" {
  type = string
  description = "Source of the root volume for the virtual machine."
  default = "https://cloud.debian.org/images/cloud/sid/daily/20241011-1897/debian-sid-generic-amd64-daily-20241011-1897.qcow2"
}