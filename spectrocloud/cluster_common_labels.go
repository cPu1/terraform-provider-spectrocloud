package spectrocloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func toAdditionalLabels(d *schema.ResourceData) map[string]string {
	tags := make(map[string]string)
	if d.Get("tags") != nil {
		for _, t := range d.Get("tags").(*schema.Set).List() {
			tag := t.(string)
			if strings.Contains(tag, ":") {
				tags[strings.Split(tag, ":")[0]] = strings.Split(tag, ":")[1]
			} else {
				tags[tag] = "spectro__tag"
			}
		}
	}
	return tags
}