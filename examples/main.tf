provider "equinix" {
    debug = true
}

resource "equinix_ecx_connection" "aws" {
    primary_name = "TERRAFORM_AWS_PRI"
    speed = 50
    speed_unit = "MB"
    seller_profile_uuid = "69ee618d-be52-468d-bc99-00566f2dd2b9" // AWS
    primary_port_uuid = "66284add-49a3-9a30-b4e0-30ac094f8af1" // EQUINIX_SVC-LD5-CX-PRI-01
    authorization_key = "111111111111" // AWS Account Number
    notifications = ["your@email.com"]
    primary_vlan_stag = 3020
    seller_metro_code = "LD"
    seller_region = "eu-west-1"
}

resource "equinix_ecx_connection" "azure" {
    speed = 50
    speed_unit = "MB"
    seller_profile_uuid = "a1390b22-bbe0-4e93-ad37-85beef9d254d" // Azure

    primary_name = "TERRAFORM_AZURE_PRI"
    primary_port_uuid = "66284add-49a3-9a30-b4e0-30ac094f8af1" // EQUINIX_SVC-LD5-CX-PRI-01
    primary_vlan_stag = 3030

    secondary_name = "TERRAFORM_AZURE_SEC"
    secondary_port_uuid = "66284add-49a5-9a50-b4e0-30ac094f8af1" // EQUINIX_SVC-LD4-CX-SEC-01
    secondary_vlan_stag = 3032

    authorization_key = "c888f96c-7f5f-4127-b339-565a89842843" // Azure Service Key (Existing Express Route)
    notifications = ["your@email.com"]
        
    seller_metro_code = "LD"
    seller_region = "westeurope"
    named_tag = "Public"
}

resource "equinix_ecx_connection" "google" {
    primary_name = "CMA_GCP_TERRAFORM"
    speed = 50
    speed_unit = "MB"
    seller_profile_uuid = "bd4570e2-d792-4a00-87f5-3bde040cdcd7" // Google Cloud Partner Interconnect Zone 1
    primary_port_uuid = "66284add-49a3-9a30-b4e0-30ac094f8af1" // EQUINIX_SVC-LD5-CX-PRI-01
    authorization_key = "0418d78e-9a30-b4e0-49a3-6e2932016439/europe-west2/1" // VLAN Pairing Key
    notifications = ["your@email.com"]
    primary_vlan_stag = 3060
    seller_metro_code = "LD"
    seller_region = "europe-west2"
}
