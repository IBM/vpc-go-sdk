[![Build Status](https://travis-ci.com/IBM/vpc-go-sdk.svg?branch=master)](https://travis-ci.com/IBM/vpc-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/vpc-go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud VPC Go SDK Version 0.0.0
Go client library to interact with the various [IBM Cloud VPC Services APIs](https://cloud.ibm.com/apidocs?category=vpc).

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
- [Setting up VPC service](#setting-up-vpc-service)
- [Setting up VPC on Classic service](#setting-up-vpc-on-classic-service)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud VPC Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name
--- | ---
[VPC](https://cloud.ibm.com/apidocs/vpc) | vpcv1
[VPC Gen 1](https://cloud.ibm.com/apidocs/vpc-on-classic) | vpcclassicv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
There are a few different ways to download and install the VPC Go SDK services for use by your
Go application:

#### `go get` command
Use this command to download and install the VPC Classic Go SDK service to allow your Go application to
use it:

```
go get -u github.com/IBM/vpc-go-sdk/vpcv1
```

To install VPC Classic Go SDK service, use the following.

```
go get -u github.com/IBM/vpc-go-sdk/vpcclassicv1
```

#### Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:


```go
import (
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.


#### `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/IBM/vpc-go-sdk/vpcclassicv1"
  version = "0.0.0"
[[constraint]]
  name = "github.com/IBM/vpc-go-sdk/vpcv1"
  version = "0.0.0"
```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Setting up VPC service

A quick example to get you up and running with VPC Go SDK service in Dallas (us-south) region.

For other regions, Refer [API Endpoints for VPC](https://cloud.ibm.com/apidocs/vpc#api-endpoint)  and update the `URL` variable accordingly.


Refer to the [VPC Release Notes](https://cloud.ibm.com/docs/vpc?topic=vpc-release-notes) document to find out latest version release.

```go
package main

import (
	"log"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

var IBMCLOUD_API_KEY = "YOUR_KEY_HERE"  // required, Add a valid API key here

func main() {
	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: IBMCLOUD_API_KEY,
	}

	// Create the service options struct.
	options := &vpcv1.VpcV1Options{
		Authenticator: authenticator,
	}

	// Instantiate the service.
	vpcService, vpcServiceErr := vpcv1.NewVpcV1(options)

	if vpcServiceErr != nil {
		log.Fatalf("Error creating VPC Service.")
	}

	// Retrieve the list of regions for your account.
	listRegionsOptions := &vpcv1.ListRegionsOptions{}
	regions, detailedResponse, err := vpcService.ListRegions(listRegionsOptions)
	if err != nil {
		log.Fatalf("Failed to list the regions: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("Regions: %+v", regions)

	// Retrieve the list of vpcs for your account.
	listVpcsOptions := &vpcv1.ListVpcsOptions{}
	vpcs, detailedResponse, err := vpcService.ListVpcs(listVpcsOptions)
	if err != nil {
		log.Fatalf("Failed to list vpcs: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("VPCs: %+v", vpcs)

	// Create an SSH key
	sshKeyOptions := &vpcv1.CreateKeyOptions{}
	sshKeyOptions.SetName("my-ssh-key")
	sshKeyOptions.SetPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDsnrSAe8eBi8mS576Z96UtYgUzDR9Sbw/s1ELxsa1KUK82JQ0Ejmz31N6sHyiT/l5533JgGL6rKamLFziMY2VX2bdyuF5YzyHhmapT+e21kuTatB50UsXzxlYEWpCmFdnd4LhwFn6AycJWOV0k3e0ePpVxgHc+pVfE89322cbmfuppeHxvxc+KSzQNYC59A+A2vhucbuWppyL3EIF4YgLwOr5iDISm1IR0+EEL3yJQIG4M2WKu526anI85QBcIWyFwQXOpdcX2eZRcd6WW2EgAM3fIOaezkm0CFrsz8rQ0MPYZI4BS2CWwg5d4Bj7SU2sjXz62gfQkQGTYWSqhizVb root@localhost")
	key, detailedResponse, err := vpcService.CreateKey(sshKeyOptions)
	if err != nil {
		log.Fatalf("Failed to create the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s created with ID: %s", *key.Name, *key.ID)

	// Delete SSH key
	deleteKeyOptions := &vpcv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(*key.ID)
	detailedResponse, err = vpcService.DeleteKey(deleteKeyOptions)
	if err != nil {
		log.Fatalf("Failed to delete the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s deleted with ID: %s", *key.Name, *key.ID)
}
```

## Setting up VPC on Classic service

A quick example to get you up and running with VPC on Classic Go SDK service in Dallas (us-south) region.

For other regions, Refer [API Endpoints for VPC on Classic](https://cloud.ibm.com/apidocs/vpc-on-classic#api-endpoint) and update the `URL` variable accordingly.

Refer to the [VPC on Classic Release Notes](https://cloud.ibm.com/docs/vpc-on-classic?topic=vpc-on-classic-release-notes) document to find out latest version release.

```go
package main

import (
	"log"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
)

var IBMCLOUD_API_KEY = "YOUR_KEY_HERE" 	// required, Add a valid API key here

func main() {
	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: IBMCLOUD_API_KEY,
	}

	// Create the service options struct.
	options := &vpcclassicv1.VpcClassicV1Options{
		Authenticator: authenticator,
	}

	// Instantiate the service.
	vpcService, vpcServiceErr := vpcclassicv1.NewVpcClassicV1(options)

	if vpcServiceErr != nil {
		log.Fatalf("Error creating VPC Service.")
	}

	// Retrieve the list of regions for your account.
	listRegionsOptions := &vpcclassicv1.ListRegionsOptions{}
	regions, detailedResponse, err := vpcService.ListRegions(listRegionsOptions)
	if err != nil {
		log.Fatalf("Failed to list the regions: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("Regions: %+v", regions)

	// Retrieve the list of vpcs for your account.
	listVpcsOptions := &vpcclassicv1.ListVpcsOptions{}
	vpcs, detailedResponse, err := vpcService.ListVpcs(listVpcsOptions)
	if err != nil {
		log.Fatalf("Failed to list vpcs: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("VPCs: %+v", vpcs)

	// Create an SSH key
	sshKeyOptions := &vpcclassicv1.CreateKeyOptions{}
	sshKeyOptions.SetName("my-ssh-key")
	sshKeyOptions.SetName("my-ssh-key")
	sshKeyOptions.SetPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDsnrSAe8eBi8mS576Z96UtYgUzDR9Sbw/s1ELxsa1KUK82JQ0Ejmz31N6sHyiT/l5533JgGL6rKamLFziMY2VX2bdyuF5YzyHhmapT+e21kuTatB50UsXzxlYEWpCmFdnd4LhwFn6AycJWOV0k3e0ePpVxgHc+pVfE89322cbmfuppeHxvxc+KSzQNYC59A+A2vhucbuWppyL3EIF4YgLwOr5iDISm1IR0+EEL3yJQIG4M2WKu526anI85QBcIWyFwQXOpdcX2eZRcd6WW2EgAM3fIOaezkm0CFrsz8rQ0MPYZI4BS2CWwg5d4Bj7SU2sjXz62gfQkQGTYWSqhizVb root@localhost")
	key, detailedResponse, err := vpcService.CreateKey(sshKeyOptions)
	if err != nil {
		log.Fatalf("Failed to create the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s created with ID: %s", *key.Name, *key.ID)

	// Delete SSH key
	deleteKeyOptions := &vpcclassicv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(*key.ID)
	detailedResponse, err = vpcService.DeleteKey(deleteKeyOptions)
	if err != nil {
		log.Fatalf("Failed to delete the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s deleted with ID: %s", *key.Name, *key.ID)

}
```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](<github-repo-url>/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).