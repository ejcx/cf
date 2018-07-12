command {
    name = "zone"
    description = <<EOF
This is a meaty description of the zone api.
EOF
    options {
      "debug" = "Turn on debug logging"
    }
    subcommands = [
      "list-zones"
    ]
}

command {
  name = "list-zones"
  description = <<EOF
This is a meaty description of the list-zones
EOF
  options {}
}

