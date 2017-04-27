package qcloud

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceQcloudClb() *schema.Resource {
	return &schema.Resource{
		Create: resourceQcloudClbCreate,
		Read:   resourceQcloudClbRead,
		Update: resourceQcloudClbUpdate,
		Delete: resourceQcloudClbDelete,

		Schema: map[string]*schema.Schema{
			"load_balancer_type": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"forward": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: true,
			},

			"load_balancer_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"domain_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"project_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}

}

func resourceQcloudClbCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	clbArgs := &CreateLoadBalancerArgs{}

	if v, ok := d.GetOk("load_balancer_type"); ok {
		clbArgs.LoadBalancerType = v.(int)

	}

	if v, ok := d.GetOk("forward"); ok {
		clbArgs.Forward = v.(int)

	}

	if v, ok := d.GetOk("load_balancer_name"); ok {
		clbArgs.LoadBalancerName = v.(string)

	}

	if v, ok := d.GetOk("domain_prefix"); ok {
		clbArgs.DomainPrefix = v.(string)

	}

	if v, ok := d.GetOk("vpc_id"); ok {
		clbArgs.VpcId = v.(string)

	}

	if v, ok := d.GetOk("subnet_id"); ok {
		clbArgs.SubnetId = v.(string)

	}

	if v, ok := d.GetOk("project_id"); ok {
		clbArgs.ProjectId = v.(int)

	}

	clb, err := client.CreateLoadBalancer(clbArgs)
	if err != nil {
		return err
	}

	d.SetId(clb.UnLoadBalancerIds[clb.DealIds[0]][0])

	return err
}

func resourceQcloudClbRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	clb, err := client.DescribeLoadBalancer(d.Id())
	if err != nil {
		d.SetId("")
		return err

	}

	if clb == nil {
		d.SetId("")
		return nil
	}

	d.Set("load_balancer_type", clb.LoadBalancerSet[0].LoadBalancerType)
	d.Set("forward", clb.LoadBalancerSet[0].LoadBalancerName)
	d.Set("load_balancer_name", clb.LoadBalancerSet[0].LoadBalancerName)
	d.Set("domain_prefix", clb.LoadBalancerSet[0].Domain)
	d.Set("vpc_id", clb.LoadBalancerSet[0].VpcId)
	d.Set("subnet_id", clb.LoadBalancerSet[0].SubnetId)
	d.Set("project_id", clb.LoadBalancerSet[0].ProjectId)

	return nil
}

func resourceQcloudClbDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	_, err := client.DeleteLoadBalancer(d.Id())
	return err
}

func resourceQcloudClbUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	d.Partial(true)

	if d.HasChange("load_balancer_name") {
		clbArgs := &ModifyLoadBalancerAttributesArgs{}
		clbArgs.LoadBalancerId = d.Id()
		clbArgs.LoadBalancerName = d.Get("load_balancer_name").(string)

		_, err := client.ModifyLoadBalancerAttributes(clbArgs)
		if err != nil {
			return err
		}

		d.SetPartial("load_balancer_name")
	}

	d.Partial(false)

	return resourceQcloudClbRead(d, meta)
}
