package equinix

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceConnection represents a layer2 Equinix Connection
func resourceConnection() *schema.Resource {
	return &schema.Resource{
		Create:        resourceConnectionCreate,
		Read:          resourceConnectionRead,
		Update:        resourceConnectionUpdate,
		Delete:        resourceConnectionDelete,
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"named_tag": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"notifications": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeSet,
				},
			},
			"authorization_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"primary_port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"primary_vlan_ctag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_vlan_stag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_zside_port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_zside_vlan_ctag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_zside_vlan_stag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"purchase_order_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_vlan_ctag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_vlan_stag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"secondary_zside_port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_zside_vlan_ctag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"profile_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seller_metro_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seller_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"speed_unit": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConnectionCreate(d *schema.ResourceData, m interface{}) error {
	//uuid := d.Get("uuid").(string)

	//d.SetId(uuid)
	client := m.(*EquinixClient).ECXConnectionsAPI

	params := client.NewCreateL2ConnectionParams()

	params.AuthorizationKey = d.Get("authorization_key").(string)
	params.NamedTag = d.Get("named_tag").(string)

	// fix type assertion
	//params.Notifications = d.Get("notifications").(string)
	params.PrimaryName = d.Get("primary_name").(string)
	params.PrimaryPortUUID = d.Get("primary_port_uuid").(string)
	params.PrimaryVlanSTag = d.Get("primary_vlan_stag").(int64)
	params.PrimaryVlanCTag = d.Get("primary_vlan_ctag").(string)

	params.PrimaryZSidePortUUID = d.Get("primary_zside_port_uuid").(string)
	params.PrimaryZSideVlanCTag = d.Get("primary_zside_vlan_ctag").(int64)
	params.PrimaryZSideVlanSTag = d.Get("primary_zside_vlan_stag").(int64)

	params.ProfileUUID = d.Get("profile_uuid").(string)
	params.PurchaseOrderNumber = d.Get("purchase_order_number").(string)

	params.SecondaryName = d.Get("secondary_name").(string)
	params.SecondaryPortUUID = d.Get("secondary_port_uuid").(string)
	params.SecondaryVlanCTag = d.Get("secondary_vlan_ctag").(string)
	params.SecondaryVlanSTag = d.Get("secondary_vlan_stag").(int64)
	params.SecondaryZSidePortUUID = d.Get("secondary_zside_port_uuid").(string)
	params.SecondaryZSideVlanCTag = d.Get("secondary_zside_vlan_ctag").(int64)
	params.SecondaryZSideVlanSTag = d.Get("secondary_zside_vlan_stag").(int64)
	params.SellerMetroCode = d.Get("seller_metro_code").(string)
	params.SellerRegion = d.Get("seller_region").(string)

	params.Speed = d.Get("speed").(int64)
	params.SpeedUnit = d.Get("speed_unit").(string)

	conn, err := client.CreateL2Connection(params)
	if err != nil {
		return err
		/**switch t := err.(type) {
		case *m.(*EquinixClient).client.CreateConnectionUsingPOSTBadRequest:
			for _, er := range t.Payload {
				fmt.Printf("Error %s with message %s\n", er.ErrorCode, er.ErrorMessage)
			}
		default:
			fmt.Printf("Error creating connection: %s\n", err.Error())
		}
		os.Exit(1)
		return errors.New("Error creating connection: %s\n", err.Error())
		**/
	}
	d.SetId(conn.Payload.PrimaryConnectionID)
	return resourceConnectionRead(d, m)
}

func resourceConnectionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*EquinixClient).ECXConnectionsAPI

	conn, err := client.GetByUUID(d.Id())
	if err != nil {
		return err
	}
	if conn {

		params.AuthorizationKey = d.Get("authorization_key").(string)
		params.NamedTag = d.Get("named_tag").(string)

		// fix type assertion
		//params.Notifications = d.Get("notifications").(string)
		d.Set("primary_name", conn.Payload.Name)

		params.PrimaryPortUUID = d.Get("primary_port_uuid").(string)
		params.PrimaryVlanSTag = d.Get("primary_vlan_stag").(int64)
		params.PrimaryVlanCTag = d.Get("primary_vlan_ctag").(string)

		params.PrimaryZSidePortUUID = d.Get("primary_zside_port_uuid").(string)
		params.PrimaryZSideVlanCTag = d.Get("primary_zside_vlan_ctag").(int64)
		params.PrimaryZSideVlanSTag = d.Get("primary_zside_vlan_stag").(int64)

		params.ProfileUUID = d.Get("profile_uuid").(string)
		params.PurchaseOrderNumber = d.Get("purchase_order_number").(string)

		params.SecondaryName = d.Get("secondary_name").(string)
		params.SecondaryPortUUID = d.Get("secondary_port_uuid").(string)
		params.SecondaryVlanCTag = d.Get("secondary_vlan_ctag").(string)
		params.SecondaryVlanSTag = d.Get("secondary_vlan_stag").(int64)
		params.SecondaryZSidePortUUID = d.Get("secondary_zside_port_uuid").(string)
		params.SecondaryZSideVlanCTag = d.Get("secondary_zside_vlan_ctag").(int64)
		params.SecondaryZSideVlanSTag = d.Get("secondary_zside_vlan_stag").(int64)
		params.SellerMetroCode = d.Get("seller_metro_code").(string)
		params.SellerRegion = d.Get("seller_region").(string)

		params.Speed = d.Get("speed").(int64)
		params.SpeedUnit = d.Get("speed_unit").(string)

		d.Set("uuid", conn.UUID)
		d.Set("named_tag", conn.NamedTag)
		d.Set("primary_name", conn.PrimaryName)
		d.Set("primary_port_uuid", conn.PortUUID)
		d.SetId(conn.UUID)
	}

	return nil
}

func resourceConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceConnectionRead(d, m)
}

func resourceConnectionDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
