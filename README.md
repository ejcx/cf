<div align="center">
  <br><br><br><br><br>
  <img src="https://raw.githubusercontent.com/ejcx/cf/master/logo.png" alt="cf logo" width="400">
  <br><br><br><br><br><br><br><br>
</div>
# cf
| The cloudflare command line interface
# Installation
```
go get -u github.com/ejcx/cf
```
# Credentials
Your cloudflare api credentials can be stored in `~/.cf/credentials` or as environment variables.
### ~/.cf/credentials
The format of the file should be like so:
```
{
  "Email": "your@email.com",
  "UserServiceKey": "v1.0-...aaa",
  "Key": "20...973"
}
```
### Environment Variables
The three environment variables that should be set are `CF_API_KEY`, `CF_API_EMAIL`, and `CF_USER_SERVICE_KEY`. If any of the environment variables are set then the credentials file is ignored
### Keychain
Coming soon...
# Usage
**The cf is a full feature cli. All product areas are divided in to commands**
```
e :) ./cf            
A CLI for interacting with Cloudflare's V4 API

Usage:
  cf [command]

Available Commands:
  cache        Commands for interacting with caching and railgun APIs
  dns          Create, read, update, and delete your dns records.
  firewall     Commands to interact with products that block access to your site
  help         Help about any command
  loadbalancer Manage, create, and describe your loadbalancers, pools, and monitor
  organization Interact with the organizations you own and have access to
  pagerule     Change how cloudflare works on a URL or subdomain basis.
  ratelimit    Configure, create, and view your zone's ratelimits.
  ssl          Control and insight in to your zone's SSL stack 
  user         Interacting with your Cloudflare account
  zone         Interact with cloudflare zones

Flags:
  -h, --help   help for cf

Use "cf [command] --help" for more information about a command.
```

**Each command has a list of subcommands that will allow you to interact with that product area.**
```
e :) cf user
  The User subcommand will help you interact with your Cloudflare account.

Usage:
  cf user [command]

Available Commands:
  billing-profile         Get the billing profile
  create-user-access-rule Create a user access rule
  delete-user-access-rule Delete a user access rule
  details                 Get user specific settings
  edit-user               Edit user account details
  list-user-access-rules  List User Access Rules
  update-user-access-rule Create a user access rule

Flags:
  -h, --help   help for user

Use "cf user [command] --help" for more information about a command.
```

**Here I am viewing the optional flags to the `edit-user` command.**
```
e :) cf user edit-user --help
Edit user account details

Usage:
  cf user edit-user [flags]

Flags:
      --country string      The country in which the user lives, example US. max length 30
      --first-name string   User's first name, max length 60
  -h, --help                help for edit-user
      --last-name string    User's last name, max length 60
      --telephone string    User's telephone number, max length 20
      --zipcode string      The zip code or postal code in which the user lives, max length 20
```


**We can easily update the first name associated with the account by specifying a new first name.**
```
e :) cf user edit-user --first-name Pete --last-name ZaHut
{
     "id": "64...4d",
     "email": "evan@xxx",
     "first_name": "Pete",
     "last_name": "ZaHut",
     "username": "e",
     "telephone": "1-867-5309",
     "zipcode": "94108",
     "created_on": "2014-10-30T02:06:57.665472Z",
     "modified_on": "2018-07-20T16:56:00.74237Z",
     "betas": [
         "dnssec_beta",
         "chat_beta",
         "stream_beta"
     ]
 }

```