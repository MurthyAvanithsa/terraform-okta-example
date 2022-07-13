locals {
  groups = csvdecode(file("csv/groups.csv"))
}


locals {
  users = csvdecode(file("csv/users.csv"))
}
