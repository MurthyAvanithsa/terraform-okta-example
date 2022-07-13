resource "okta_group_memberships" "sprout_groups_all_staff_membership" {
  group_id = okta_group.sprout_groups["group_all_staff"].id
  users = [
    okta_user.sprout_users["ajay_avanthsa@oktaice.com"].id,
    okta_user.sprout_users["anand_kolli@oktaice.com"].id,
  ]
  depends_on = [
    okta_user.sprout_users, okta_group.sprout_groups
  ]
}
