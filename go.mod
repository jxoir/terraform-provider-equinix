module github.com/jxoir/terraform-provider-equinix

require (
	github.com/hashicorp/go-hclog v0.7.0 // indirect
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/terraform v0.12.0
	github.com/jxoir/equinix-tools v0.1.1
	github.com/jxoir/go-ecxfabric v0.0.0-20181101112837-9ea2dc638437
)

replace github.com/jxoir/equinix-tools v0.1.1 => ../equinix-tools
