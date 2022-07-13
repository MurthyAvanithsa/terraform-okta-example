
resource "okta_user" "sprout_users" {

  for_each   = { for inst in local.users : inst.user_login => inst }
  first_name = each.value.first_name
  last_name  = each.value.last_name
  login      = each.value.user_login
  email      = each.value.email

}

