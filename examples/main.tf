provider "equinix" {
    debug = true
}

resource "equinix_ecx_connection" "aws" {
    primary_name = "TERRAFORM_AWS_TEST"
    speed = 50
    speed_unit = "MB"
    seller_profile_uuid = "69ee618d-be52-468d-bc99-00566f2dd2b9" // AWS
    primary_port_uuid = "66284add-49a3-9a30-b4e0-30ac094f8af1" // AWS DX "Standard"
    authorization_key = "123456789012" // AWS Account Number
    notifications = ["some@email.com"]
    primary_vlan_stag = 3020
    seller_metro_code = "LD"
    seller_region = "eu-west-1"
}
