package equinix

import "github.com/hashicorp/terraform/helper/schema"

func resourceConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectionCreate,
		Read:   resourceConnectionRead,
		Update: resourceConnectionUpdate,
		Delete: resourceConnectionDelete,

		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConnectionCreate(d *schema.ResourceData, m interface{}) error {
	uuid := d.Get("uuid").(string)
	d.SetId(uuid)
	return resourceConnectionRead(d, m)
}

func resourceConnectionRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceConnectionRead(d, m)
}

func resourceConnectionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
