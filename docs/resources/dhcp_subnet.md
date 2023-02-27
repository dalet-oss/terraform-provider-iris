---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iris_dhcp_subnet Resource - terraform-provider-iris"
subcategory: ""
description: |-
  
---

# iris_dhcp_subnet (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `pool` (Block List, Min: 1) (see [below for nested schema](#nestedblock--pool))

### Optional

- `cidr` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--pool"></a>
### Nested Schema for `pool`

Required:

- `first` (String)
- `last` (String)

