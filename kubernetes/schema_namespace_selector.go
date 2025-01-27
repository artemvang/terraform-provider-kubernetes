package kubernetes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func namespaceSelectorFields(updatable bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"match_expressions": {
			Type:        schema.TypeList,
			Description: "A list of namespace selector requirements. The requirements are ANDed.",
			Optional:    true,
			ForceNew:    !updatable,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key": {
						Type:        schema.TypeString,
						Description: "The namespace key that the selector applies to.",
						Optional:    true,
						ForceNew:    !updatable,
					},
					"operator": {
						Type:        schema.TypeString,
						Description: "A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.",
						Optional:    true,
						ForceNew:    !updatable,
					},
					"values": {
						Type:        schema.TypeSet,
						Description: "An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.",
						Optional:    true,
						ForceNew:    !updatable,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Set:         schema.HashString,
					},
				},
			},
		},
		"match_labels": {
			Type:        schema.TypeMap,
			Description: "A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
			Optional:    true,
			ForceNew:    !updatable,
		},
	}
}
