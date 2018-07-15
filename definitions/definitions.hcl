command {
    name = "zone"
    description = <<EOF
This is a meaty description of the zone api.
EOF
    shortdescription = "Commands for interacting zones"
    options {
      "debug" = "Turn on debug logging"
    }
    subcommands = [
      "list-zones"
    ]
    toplevel = true
}

command {
  name = "list-zones"
  shortdescription = "Command for listing zones"
  description = <<EOF
This is a meaty description of the list-zones
EOF
  options {}
}

