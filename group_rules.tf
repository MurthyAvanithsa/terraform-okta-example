resource "okta_group_rule" "rule_0prrnzdllC5mY7bz01d6" {
  name              = "Add to AllAdmin"
  expression_type   = "urn:okta:expression:1.0"
  expression_value  = "isMemberOfAnyGroup(\"${okta_group.sprout_groups["group_all_admin"].id}\")"
  group_assignments = [okta_group.sprout_groups["group_all_staff"].id]
  depends_on        = [okta_group.sprout_groups]
}
