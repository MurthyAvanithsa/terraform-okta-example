resource "okta_policy_signon" "P101" {
  name        = "p101"
  status      = "ACTIVE"
  description = "p101"
  depends_on = [
    okta_group.sprout_groups
  ]
  groups_included = [okta_group.sprout_groups["group_all_staff"].id]
}
