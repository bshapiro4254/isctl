![Build](https://github.com/cgascoig/isctl/workflows/Build/badge.svg)

# isctl - CLI for Cisco Intersight
`isctl` is a `kubectl`-insipired CLI for the Cisco Intersight service. 

## Features

* Generated automatically from the Intersight OpenAPI v3 spec - it should be easy to keep up with new Intersight features. 
* Written in Go and distributed as MacOS, Linux and Windows binaries - should work anywhere
* Human, JSON or YAML output
* [JSONpath](https://goessner.net/articles/JsonPath/) support for extraction and transformation of the data returned from the API


# Installation

## MacOS

If you use [Homebrew](https://brew.sh), the easiest way to install `isctl` is:

```
brew install cgascoig/isctl/isctl
```

If you don't use Homebew:

* Download the latest release from the [Releases](releases/) page. 
* Unzip and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 


## Windows

* Download the latest release from the [Releases](releases/) page. 
* Unzip and move the `isctl.exe` binary somewhere that is on your path. 

## Linux

* Download the latest release from the [Releases](releases/) page. 
* Extract the `.tar.gz` and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 


# Usage

## Initial configuration

### Obtain your API Key

`isctl` interacts with the Cisco Intersight REST API, so it needs an API key. 

1. Login to the Intersight GUI. 
2. Generate a new API Key (under Settings -> API Keys). Choose "API key for OpenAPI schema version 2" as the API Key Purpose. 
3. Save the key somewhere on your desktop and make a note of the key ID. 

### Configure `isctl`

Run `isctl configure` to configure it to use your API key. Follow the prompts for your key ID and path to the key file. (Note: the path to the key file should be an absolute path and doesn't support shell metacharacters such as "`~`")

```
keyID is currently ''
Enter new keyID or press Enter to keep existing: 123456789012345678901234/123456789012345678901234/123456789012345678901234
key filename is currently ''
Enter new key file name or press Enter to keep existing: /Users/chris/intersight-key.pem
2020/06/25 17:28:22 Writing config file
```

## Querying Intersight

To retrieve resources from Intersight, use the `isctl get ...` commands. 

For example, to query all objects of a certain type:

```
isctl get ntp policy
```
Output:
```
------------  -----------------------------  ------------------------  -----------------------------
    Enabled                           Moid                      Name                     NtpServers
------------  -----------------------------  ------------------------  -----------------------------
      False       5ecb069c6275722d3143e668           test-ntp-policy

       True       5ee1aa076275722d3122a944          cg-tf-ntp-test-1       10.10.10.10, 10.10.10.12
------------  -----------------------------  ------------------------  -----------------------------
```
Notice that all objects are returned in a human-readable table. 

You can query a single object based on its MoId:

```
isctl get ntp policy moid 5ee1aa076275722d3122a944
```

Or its name:
```
isctl get ntp policy --name cg-tf-ntp-test-1
```

Output:
```
2020/06/25 17:40:04 Single result, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.
Enabled: true
Moid: 5ee1aa076275722d3122a944
Name: cg-tf-ntp-test-1
NtpServers:
- 10.10.10.10
- 10.10.10.12
```
Notice that when displaying a single object, table format is not used. 

The default outputs above automatically hide many of the boilerplate attributes of the objects (e.g. `ClassId`, `Organization`, etc.). If you specify JSON or YAML as the output format all attributes will be included. For example:

```
isctl get ntp policy moid 5ee1aa076275722d3122a944 --output json
```
Output:
```json
{
  "AccountMoid": "123456789012345678901234",
  "Ancestors": [],
  "ClassId": "ntp.Policy",
  "CreateTime": "2020-06-11T03:50:32.062Z",
  "Description": "",
  "DomainGroupMoid": "123456789012345678901234",
  "Enabled": true,
  "ModTime": "2020-06-11T03:50:32.063Z",
  "Moid": "5ee1aa076275722d3122a944",
  "Name": "cg-tf-ntp-test-1",
  "NtpServers": [
    "10.10.10.10",
    "10.10.10.12"
  ],
  "ObjectType": "ntp.Policy",
  "Organization": {
    "ClassId": "mo.MoRef",
    "Moid": "123456789012345678901234",
    "ObjectType": "organization.Organization",
    "link": "https://www.intersight.com/api/v1/organization/Organizations/123456789012345678901234"
  },
  "Owners": [
    "123456789012345678901234"
  ],
  "PermissionResources": [
    {
      "ClassId": "mo.MoRef",
      "Moid": "123456789012345678901234",
      "ObjectType": "organization.Organization",
      "link": "https://www.intersight.com/api/v1/organization/Organizations/123456789012345678901234"
    }
  ],
  "Profiles": [],
  "SharedScope": "",
  "Tags": []
}
```

You can use JSONPath to filter/transform the returned data - to output just the name attribute of all the returned objects:

```
isctl get ntp policy -o yaml --jsonpath '$[*].Name'
```
Output:
```
- test-ntp-policy
- cg-tf-ntp-test-1
```

Or even to get multiple attributes from the returned objects:

```
isctl get ntp policy moid 5ee1aa076275722d3122a944 --jsonpath '$["Name","Enabled"]'
```
Output:
```
cg-tf-ntp-test-1
True
```

## Creating Intersight resources

To create resources in Intersight, use the `isctl create ...` commands. 

The attributes for the new object can be passed in as command line flags:

```
isctl create ntp policy --Name isctl-test-1 --NtpServers 1.1.1.1,2.2.2.2 --Organization default
```

Output:
```
2020/07/19 13:02:13 Single result, falling back to vertical output. NOTE: this is not valid YAML; use --output yaml to get valid YAML.
Enabled: true
Moid: 5f13b7b56275722d31133e63
Name: isctl-test-1
NtpServers:
- 1.1.1.1
- 2.2.2.2
```

The output of the `create` commands is the resource after it was created in Intersight, so the Moid, etc. are included.

You can also define the attributes of the new object via standard input:

```
isctl create ntp policy --bodyformat json
Waiting for JSON body:
{
    "Name": "isctl-ntp-test",
    "Enabled": true,
    "NtpServers": [
        "1.1.1.1"
    ],
    "Organization": {
        "Moid": "123456789012345678901234",
        "ClassId": "mo.MoRef"
    }
}
```

This is useful when you have your object defined as JSON in a file, for example:

```
cat ntp_policy.json | isctl create ntp policy --bodyformat json
```

## Deleting resources

Deleting resources is quite simple if you know the Moid:

```
isctl delete ntp policy moid 123456789012345678901234
```

If you don't know the Moid, you can combine two commands together, for example this command uses a `get` to find the object by name, then uses the JSONPath query to extract the Moid and finally calls the `delete` command:

```
isctl delete ntp policy moid $(./build/isctl get ntp policy --name isctl-test-1 --jsonpath '$.Moid')
```

## Updating resources

Updating resources is also quite simple if you know the Moid. For example, if you want to set the `Enabled` attribute in an NTP policy to `False`:

```
isctl update ntp policy moid 123456789012345678901234 --Enabled=False
```

If you don't know the Moid, use the same technique as described for deletion:

```
isctl update ntp policy moid $(./build/isctl get ntp policy --name isctl-test-1 --jsonpath '$.Moid') --Enabled=False
```

# Development

## How is this built?

There are a number of steps in the build process as the majority of the code is generated automatically:

1. [OpenAPITools OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) is used to generate the Go models and API client from the Intersight OpenAPI v3 spec. The OpenAPI Generator Go templates are customised to also generate an `operations.yaml` file. 
2. The Go program in `generator-postprocess` is used to parse the `operations.yaml` file and generate the command line structure in `cmd/cli.go`. This postprocess step also cleans up the command naming and creates the help text. 
3. Finally, the `isctl` command is built from the generated `cmd/cli.go` along with some supporting code. 

## Creating a release
The `Makefile` is used for most of the post processing and building `isctl` for the local system but [GoReleaser](https://goreleaser.com) is used for creating releases on GitHub. 

```
# Create tag and push to GitHub
git tag vX.Y.Z
git push origin --tags

# Create the release
goreleaser --rm-dist
```