---
page_title: "spectrocloud_cluster_edge_native Resource - terraform-provider-spectrocloud"
subcategory: ""
description: |-
  
---

# Resource `spectrocloud_cluster_edge_native`





## Schema

### Required

- **cloud_config** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--cloud_config))
- **machine_pool** (Block Set, Min: 1) (see [below for nested schema](#nestedblock--machine_pool))
- **name** (String)

### Optional

- **apply_setting** (String)
- **backup_policy** (Block List, Max: 1) (see [below for nested schema](#nestedblock--backup_policy))
- **cloud_account_id** (String)
- **cluster_profile** (Block List) (see [below for nested schema](#nestedblock--cluster_profile))
- **cluster_rbac_binding** (Block List) (see [below for nested schema](#nestedblock--cluster_rbac_binding))
- **host_config** (Block List) (see [below for nested schema](#nestedblock--host_config))
- **id** (String) The ID of this resource.
- **namespaces** (Block List) (see [below for nested schema](#nestedblock--namespaces))
- **os_patch_after** (String)
- **os_patch_on_boot** (Boolean)
- **os_patch_schedule** (String)
- **pack** (Block List) (see [below for nested schema](#nestedblock--pack))
- **scan_policy** (Block List, Max: 1) (see [below for nested schema](#nestedblock--scan_policy))
- **skip_completion** (Boolean)
- **tags** (Set of String)
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-only

- **cloud_config_id** (String)
- **kubeconfig** (String)

<a id="nestedblock--cloud_config"></a>
### Nested Schema for `cloud_config`

Required:

- **ssh_key** (String)
- **vip** (String)

Optional:

- **ntp_servers** (Set of String)


<a id="nestedblock--machine_pool"></a>
### Nested Schema for `machine_pool`

Required:

- **name** (String)

Optional:

- **additional_labels** (Map of String)
- **control_plane** (Boolean)
- **control_plane_as_worker** (Boolean)
- **host_uids** (List of String)
- **taints** (Block List) (see [below for nested schema](#nestedblock--machine_pool--taints))
- **update_strategy** (String)

<a id="nestedblock--machine_pool--taints"></a>
### Nested Schema for `machine_pool.taints`

Required:

- **effect** (String)
- **key** (String)
- **value** (String)



<a id="nestedblock--backup_policy"></a>
### Nested Schema for `backup_policy`

Required:

- **backup_location_id** (String)
- **expiry_in_hour** (Number)
- **prefix** (String)
- **schedule** (String)

Optional:

- **include_cluster_resources** (Boolean)
- **include_disks** (Boolean)
- **namespaces** (Set of String)


<a id="nestedblock--cluster_profile"></a>
### Nested Schema for `cluster_profile`

Required:

- **id** (String) The ID of this resource.

Optional:

- **pack** (Block List) (see [below for nested schema](#nestedblock--cluster_profile--pack))
- **type** (String)

<a id="nestedblock--cluster_profile--pack"></a>
### Nested Schema for `cluster_profile.pack`

Required:

- **name** (String)
- **tag** (String)
- **values** (String)

Optional:

- **manifest** (Block List) (see [below for nested schema](#nestedblock--cluster_profile--pack--manifest))
- **registry_uid** (String)
- **type** (String)

<a id="nestedblock--cluster_profile--pack--manifest"></a>
### Nested Schema for `cluster_profile.pack.manifest`

Required:

- **content** (String)
- **name** (String)




<a id="nestedblock--cluster_rbac_binding"></a>
### Nested Schema for `cluster_rbac_binding`

Required:

- **type** (String)

Optional:

- **namespace** (String)
- **role** (Map of String)
- **subjects** (Block List) (see [below for nested schema](#nestedblock--cluster_rbac_binding--subjects))

<a id="nestedblock--cluster_rbac_binding--subjects"></a>
### Nested Schema for `cluster_rbac_binding.subjects`

Required:

- **name** (String)
- **type** (String)

Optional:

- **namespace** (String)



<a id="nestedblock--host_config"></a>
### Nested Schema for `host_config`

Optional:

- **external_traffic_policy** (String)
- **host_endpoint_type** (String)
- **ingress_host** (String)
- **load_balancer_source_ranges** (String)


<a id="nestedblock--namespaces"></a>
### Nested Schema for `namespaces`

Required:

- **name** (String)
- **resource_allocation** (Map of String)


<a id="nestedblock--pack"></a>
### Nested Schema for `pack`

Required:

- **name** (String)
- **tag** (String)
- **values** (String)

Optional:

- **registry_uid** (String)


<a id="nestedblock--scan_policy"></a>
### Nested Schema for `scan_policy`

Required:

- **configuration_scan_schedule** (String)
- **conformance_scan_schedule** (String)
- **penetration_scan_schedule** (String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String)
- **delete** (String)
- **update** (String)

