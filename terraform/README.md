# Canberra Project

This is the **Canberra** project, where we set up a Kubernetes cluster using [Kubespray](https://github.com/kubernetes-sigs/kubespray) and Ansible, deployed on virtual machines provisioned via Terraform.

The VMs are configured using Terraform in a similar way as in the [Cairo](https://github.com/tnsoftbear/study-coding/tree/main/terraform/cairo) project.

The `templates/inventory.tpl` file contains an HCL template that Terraform uses to generate the `inventory.ini` file. This is achieved with the [`local_file`](https://registry.terraform.io/providers/hashicorp/local/latest/docs/resources/file) resource from the [hashicorp/local](https://registry.terraform.io/providers/hashicorp/local/latest) provider, which takes output variables from the `masters` and `nodes` modules.

## Prerequisites

Before installing Kubernetes, you need to resize the VM disk to fit Kubernetes, as the initial qcow2 disk size is only 2GB. After installation, the master node will use around `3.5GB`. Additionally, 2GB of memory is allocated to both the master node and agent nodes, as Ansible would otherwise show a memory shortage error.

## Steps to Install

```bash
# Resize the VM disk to have more space
$ qemu-img resize /qemu/qcow2/debian-sid-generic-amd64-daily-20241011-1897.qcow2 +3G

# Clone the Kubespray repository
$ git clone https://github.com/kubernetes-sigs/kubespray.git
$ cd kubespray/

# Checkout the specific Kubespray version
$ git checkout -b v2.26.0 tags/v2.26.0

# Install virtual environment tools
$ pamac install python-virtualenvwrapper
$ virtualenv ./.env
$ . ./.env/bin/activate

# Install Kubespray dependencies
(.venv) pip install -r ./requirements.txt

# Test the connection to all nodes using Ansible
(.venv) ansible -i ./inventory.ini all --user igor -b -m ping
# Update installed packages
(.venv) ansible -i ./inventory.ini all --user igor -b -m ansible.builtin.apt -a 'name="*" state=latest update_cache=yes'

# Run from the Kubespray directory to use the ansible.cfg where roles paths are set
(.venv) cd ./kubespray
(.venv) ansible-playbook --user igor -b -i ../inventory.ini ./cluster.yml
```

## Post-installation Resource Usage

After Kubernetes is installed, here are the resource usage statistics for the master and agent nodes:

### Master Node

```sh
$ df -h
/dev/vda1       4.8G  3.5G  1.1G  77% /
$ free -h
               total        used        free      shared  buff/cache   available
Mem:           1.9Gi       1.5Gi        90Mi       553Mi       1.0Gi       456Mi
```

### Agent Node

```sh
$ df -h
/dev/vda1       4.8G  3.0G  1.7G  65% /
$ free -h
               total        used        free      shared  buff/cache   available
Mem:           1.9Gi       995Mi       103Mi       424Mi       1.4Gi       978Mi
```

## Links

* [hashicorp/local provider](https://registry.terraform.io/providers/hashicorp/local/latest)
* [github: kubespray](https://github.com/kubernetes-sigs/kubespray)

---
