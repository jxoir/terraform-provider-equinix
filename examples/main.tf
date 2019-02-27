provider "equinix" {
    debug = true
}

resource "equinix_ecx_connection" "jxo" {
	uuid = "e235ab08-09ef-40c7-a926-5943df612046"
}
