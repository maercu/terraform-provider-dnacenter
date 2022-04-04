package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePnpDeviceSiteClaim() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).
	•	Claim a device based on DNA-C Site based design process. Different parameters are required for different device
		platforms.
	
`,

		CreateContext: resourcePnpDeviceSiteClaimCreate,
		ReadContext:   resourcePnpDeviceSiteClaimRead,
		DeleteContext: resourcePnpDeviceSiteClaimDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"config_parameters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
						"device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"image_info": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"image_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"skip": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
										Required:     true,
										ForceNew:     true,
									},
								},
							},
						},
						"site_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpDeviceSiteClaimCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method 1: ClaimADeviceToASite")
	request1 := expandRequestPnpDeviceClaimToSiteClaimADeviceToASite(ctx, "parameters.0", d)

	response1, restyResp1, err := client.DeviceOnboardingPnp.ClaimADeviceToASite(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ClaimADeviceToASite", err,
			"Failure at ClaimADeviceToASite, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenDeviceOnboardingPnpClaimADeviceToASiteItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ClaimADeviceToASite response",
			err))
		return diags
	}
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourcePnpDeviceSiteClaimRead(ctx, d, m)
}

func resourcePnpDeviceSiteClaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourcePnpDeviceSiteClaimUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePnpDeviceSiteClaimRead(ctx, d, m)
}

func resourcePnpDeviceSiteClaimDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
func expandRequestPnpDeviceClaimToSiteClaimADeviceToASite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASite {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_info")))) {
		request.ImageInfo = expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteImageInfo(ctx, key+".image_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_info")))) {
		request.ConfigInfo = expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteConfigInfo(ctx, key+".config_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteImageInfo(ctx context.Context, key string, d *schema.ResourceData) dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); ok {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip")); ok {
		request.Skip = interfaceToBool(v)
	}

	return request
}

func expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteConfigInfo(ctx context.Context, key string, d *schema.ResourceData) dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); ok {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteConfigInfoConfigParametersArray(ctx, key+".config_parameters", d)
	} else {
		request.ConfigParameters = []dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters{}
	}
	return request
}

func expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteConfigInfoConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) []dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters{}
	key = fixKeyAccess(key)
	o := d.Get(key)

	objs := o.([]interface{})
	if len(objs) == 0 {
		return request
	}
	for item_no, _ := range objs {
		i := expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteConfigInfoConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)

		request = append(request, i)

	}

	return request
}

func expandRequestPnpDeviceClaimToSiteClaimADeviceToASiteConfigInfoConfigParameters(ctx context.Context, key string, d *schema.ResourceData) dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); ok {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); ok {
		request.Value = interfaceToString(v)
	}

	return request
}

func flattenDeviceOnboardingPnpClaimADeviceToASiteItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpClaimADeviceToASite) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}