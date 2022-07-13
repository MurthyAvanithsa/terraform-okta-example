

resource "okta_group" "sprout_groups" {
  for_each    = { for inst in local.groups : inst.group_id => inst }
  name        = each.value.group_name
  description = each.value.group_desc
  depends_on = [
    okta_user.sprout_users
  ]
}
