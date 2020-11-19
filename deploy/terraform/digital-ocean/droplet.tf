resource "digitalocean_droplet" "debian_backend_production" {
  name               = "debian-backend-production"
  image              = "debian-10-x64"
  size               = "s-1vcpu-1gb"
  count              = 1
  region             = "nyc1"
  private_networking = false
  backups            = true

  ssh_keys = [
    var.ssh_fingerprint
  ]

  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    agent       = true
    private_key = file("${var.pvt_key}")
    timeout     = "2m"
  }

  # Copy the script to download all dependencies
  provisioner "file" {
    source      = abspath("provisioner")
    destination = "/tmp/"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/provisioner/*.sh",
      "/tmp/provisioner/install.sh",
    ]
  }
}
