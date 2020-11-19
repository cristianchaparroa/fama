# your DigitalOcean Personal Access Token
variable "do_token" {}

# public key location, so it can be installed into new droplets
variable "pub_key" {}

# private key location, so Terraform can connect to new droplets
variable "pvt_key" {}

# fingerprint of SSH key
variable "ssh_fingerprint" {}

provider "digitalocean" {
  token = var.do_token
}