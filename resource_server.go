package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
        return &schema.Resource{
                Create: resourceServerCreate,
                Read:   resourceServerRead,
                Update: resourceServerUpdate,
                Delete: resourceServerDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
        address := d.Get("address").(string)
        d.SetId(address)
        return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
client := m.(*MyClient)

  // Attempt to read from an upstream API
  obj, ok := client.Get(d.Id())

  // If the resource does not exist, inform Terraform. We want to immediately
  // return here to prevent further processing.
  if !ok {
    d.SetId("")
    return nil
  }

  d.Set("address", obj.Address)
 return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
        // Enable partial state mode
        d.Partial(true)

        if d.HasChange("address") {
                // Try updating the address
                if err := updateAddress(d, m); err != nil {
                        return err
                }

                d.SetPartial("address")
        }

        // If we were to return here, before disabling partial mode below,
        // then only the "address" field would be saved.

        // We succeeded, disable partial mode. This causes Terraform to save
        // all fields again.
        d.Partial(false)

        return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
// d.SetId("") is automatically called assuming delete returns no errors, but
  // it is added here for explicitness.
        d.SetId("")
 return nil
}

