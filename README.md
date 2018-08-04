<div align="left">
  <img src="https://raw.githubusercontent.com/ejcx/cf/master/logo.png" alt="cf logo" width="300">
</div>

A cloudflare command line interface. It makes heavy use of [cloudflare-go](https://github.com/cloudflare/cloudflare-go)

# Installation
```
go get -u github.com/ejcx/cf
```

# Demo
<div align="left">
  <img src="https://i.imgur.com/Js5rwCC.gif" alt="cf logo" width="600">
</div>

# Usage
**The cf is a full featured cli. All product areas are divided in to subcommands. All commands and subcommands will describe themselves and list options, required arguments, and information about what the command does.**

<div align="left">
  <img src="https://i.imgur.com/1faa2zP.gif" alt="cf logo" width="600">
</div>

**Interacting with the full Cloudflare API and retrieving the most detailed information can be done quickly and easily.**

<div align="left">
  <img src="https://i.imgur.com/JZqBJSI.gif" alt="cf logo" width="600">
</div>

**Automate common operations quickly and easily**

<div align="left">
  <img src="https://i.imgur.com/nScexK0.gif" alt="cf logo" width="600">
</div>

# Credentials
Your cloudflare api credentials can be stored in `~/.cf/credentials` as
environment variables or in your keychain.

### Keychain
To store your credentials in your keychain run `cf configure` and enter your
email, apikey, and user service apikey.
```
e :) cf configure
Cloudflare Email: evan@cloudflare.com
Cloudflare APIKey:
Origin CA APIKey:

e :)
```
### Ignore Keychain
To avoid storing your credentials in the keychain, specify the `--no-keychain`
flag when calling `cf configure` to store your api credentials in plaintext.

```
e :) cf configure --help
A command for configuring your cloudflare api credentials

Usage:
  cf configure [flags]

Flags:
  -h, --help          help for configure
      --no-keychain   Do not attempt to store cloudflare api credentials in the keychain. Just use plaintext file.

e :) cf configure --no-keychain
Cloudflare Email: evan@cloudflare.com
Cloudflare APIKey: 
Service APIKey: 

e :) cat ~/.cf/credentials
{"Email":"evan@cloudflare","Key":"xxx","UserServiceKey":"yyy","Keychain":false}% 
```

### Environment Variables
The three environment variables that should be set are `CF_API_KEY`, `CF_API_EMAIL`, and `CF_USER_SERVICE_KEY`. If any of the environment variables are set then the credentials file is ignored