package equinix

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	apiconnections "github.com/jxoir/go-ecxfabric/buyer/client/connections"
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
				Optional: true,
			},
			"notifications": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					//Type: schema.TypeSet,
					//Set:  schema.HashString,
					//Elem: &schema.Schema{Type: schema.TypeString},
					Type: schema.TypeString,
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
			"redundant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundancy_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundancy_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"seller_profile_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
	if d.Get("primary_vlan_stag") != nil {
		params.PrimaryVlanSTag = int64(d.Get("primary_vlan_stag").(int))
	}
	params.PrimaryVlanCTag = d.Get("primary_vlan_ctag").(string)

	params.PrimaryZSidePortUUID = d.Get("primary_zside_port_uuid").(string)
	if d.Get("primary_zside_vlan_ctag") != nil {
		params.PrimaryZSideVlanCTag = int64(d.Get("primary_zside_vlan_ctag").(int))
	}
	if d.Get("primary_zside_vlan_stag") != nil {
		params.PrimaryZSideVlanSTag = int64(d.Get("primary_zside_vlan_stag").(int))
	}
	params.ProfileUUID = d.Get("seller_profile_uuid").(string)
	params.PurchaseOrderNumber = d.Get("purchase_order_number").(string)

	params.SecondaryName = d.Get("secondary_name").(string)
	params.SecondaryPortUUID = d.Get("secondary_port_uuid").(string)
	params.SecondaryVlanCTag = d.Get("secondary_vlan_ctag").(string)

	if d.Get("secondary_vlan_stag") != nil {
		params.SecondaryVlanSTag = int64(d.Get("secondary_vlan_stag").(int))
	}

	params.SecondaryZSidePortUUID = d.Get("secondary_zside_port_uuid").(string)

	if d.Get("secondary_zside_vlan_ctag") != nil {
		params.SecondaryZSideVlanCTag = int64(d.Get("secondary_zside_vlan_ctag").(int))
	}
	if d.Get("secondary_zside_vlan_stag") != nil {
		params.SecondaryZSideVlanSTag = int64(d.Get("secondary_zside_vlan_stag").(int))
	}

	params.SellerMetroCode = d.Get("seller_metro_code").(string)
	params.SellerRegion = d.Get("seller_region").(string)

	params.Speed = int64(d.Get("speed").(int))
	params.SpeedUnit = d.Get("speed_unit").(string)

	notifications := d.Get("notifications").(*schema.Set).List()
	if len(notifications) > 0 {
		nfs := expandStringList(notifications)
		params.Notifications = nfs
	}

	conn, err := client.CreateL2Connection(params)
	if err != nil {
		//return err
		switch t := err.(type) {
		case *apiconnections.CreateConnectionUsingPOSTBadRequest:
			for _, er := range t.Payload {
				return fmt.Errorf("Error %s: %s - %s - %s", er.ErrorCode, er.ErrorMessage, er.Property, er.MoreInfo)
			}
		default:
			return fmt.Errorf("Error creating connection: %s", err.Error())
		}
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
	if conn != nil {

		//params.AuthorizationKey = d.Get("authorization_key").(string)
		//params.NamedTag = d.Get("named_tag").(string)

		d.Set("authorization_key", conn.Payload.AuthorizationKey)
		d.Set("named_tag", conn.Payload.NamedTag)
		// fix type assertion
		//params.Notifications = d.Get("notifications").(string)
		d.Set("primary_name", conn.Payload.Name)

		d.Set("primary_port_uuid", conn.Payload.PortUUID)

		// Should fail on 32-bit systems with large ints
		d.Set("primary_vlan_stag", int(conn.Payload.VlanSTag))
		d.Set("primary_vlan_ctag", conn.Payload.ZSideVlanCTag)

		d.Set("primary_zside_vlan_ctag", conn.Payload.ZSideVlanCTag)

		d.Set("seller_profile_uuid", conn.Payload.SellerServiceUUID)
		d.Set("purchase_order_number", conn.Payload.PurchaseOrderNumber)

		d.Set("redundant_uuid", conn.Payload.RedundantUUID)
		d.Set("redundancy_type", conn.Payload.RedundancyType)
		d.Set("redundancy_group", conn.Payload.RedundancyGroup)
		d.Set("remote", conn.Payload.Remote)
		d.Set("status", conn.Payload.Status)
		/** d.Set("secondary_port_uuid", conn.Payload.SecondaryPortUUID)
		d.Set("secondary_vlan_ctag", conn.Payload.SecondaryVlanCTag)
		d.Set("secondary_vlan_stag", conn.Payload.SecondaryVlanSTag)
		d.Set("secondary_zside_port_uuid", conn.Payload.SecondaryZSidePortUUID)
		d.Set("secondary_zside_vlan_ctag", conn.Payload.SecondaryZSideVlanCTag)
		d.Set("secondary_zside_vlan_stag", conn.Payload.SecondaryZSideVlanSTag)
		**/

		d.Set("notifications", flattenStringList(conn.Payload.Notifications))
		d.Set("seller_metro_code", conn.Payload.SellerMetroCode)
		//d.Set("seller_region", conn.Payload.Seller)

		d.Set("speed", int(conn.Payload.Speed))
		d.Set("speed_unit", conn.Payload.SpeedUnit)

		d.Set("uuid", conn.Payload.UUID)

		d.SetId(conn.Payload.UUID)
	}

	return nil
}

func resourceConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	resourceConnectionDelete(d, m)
	time.Sleep(30 * time.Second)
	return resourceConnectionCreate(d, m)
}

func resourceConnectionDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*EquinixClient).ECXConnectionsAPI

	_, err := client.DeleteByUUID(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting connection: %s", err.Error())
	}

	d.SetId("")

	secondaryConn := d.Get("redundant_uuid")
	if secondaryConn != "" {
		_, err := client.DeleteByUUID(secondaryConn.(string))
		if err != nil {
			return fmt.Errorf("Error deleting secondary connection: %s", err.Error())
		}

		d.Set("redundant_uuid", "")
	}

	return nil
}
