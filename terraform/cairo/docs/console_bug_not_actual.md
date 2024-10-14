There was a bug in cloud images where they required a console to be defined: [Cloud images fail to boot when a serial port is not available](https://bugs.launchpad.net/cloud-images/+bug/1573095). Although this issue has been resolved and is no longer relevant today, I’m keeping this information for reference.

```hcl
resource "libvirt_domain" "server" {
  ...
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
}
```

---

[index](../README.md)
