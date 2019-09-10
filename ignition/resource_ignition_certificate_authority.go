package ignition

import (
	"github.com/coreos/ignition/config/v2_2/types"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCertifcateAuthority() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCerticateAuthorityExists,
		Read:   resourceCertificateAuthorityRead,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"verification_hash": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCertificateAuthorityRead(d *schema.ResourceData, meta interface{}) error {
	id, err := buildCerticateAuthority(d, globalCache)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceCerticateAuthorityExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := buildCerticateAuthority(d, globalCache)
	if err != nil {
		return false, err
	}

	return id == d.Id(), nil
}

func buildCerticateAuthority(d *schema.ResourceData, c *cache) (string, error) {
	certicateAuthority := &types.CaReference{
		Source:       d.Get("source").(string),
		Verification: types.Verification{Hash: d.Get("verification_hash").(*string)},
	}

	if err := handleReport(certicateAuthority.ValidateSource()); err != nil {
		return "", err
	}

	return c.addCerticateAuthority(certicateAuthority), nil
}
