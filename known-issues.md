# Known Issues for vpc-go-sdk

## Security Group Rules and Network ACL Rules - backward compatibility issue
#### Publication date - "2025-12-18"

### Affected component
vpc-go-sdk

### Affected operations
Security group rules and Network ACL rules 

### Issue Summary
#### Golang SDK Backward compatibility with new protocols in Network Access Controls and Security Group rules 
Following the new support for [all IPv4 protocols for ACL and Security Group rules](https://cloud.ibm.com/docs/vpc?topic=vpc-release-notes#vpc-dec1225), earlier versions of the Golang SDK must be updated to avoid the following parsing error when handling rules with the new protocols:

```
error unmarshalling vpcv1.SecurityGroupCollection: error unmarshalling property 'security_groups' as []vpcv1.SecurityGroup: error unmarshalling property 'rules' as []vpcv1.SecurityGroupRuleIntf: unrecognized value for discriminator property 'protocol': any
```

The patched SDKs implement the correct fallback behavior and error identifiers with the correct model name (e.g. NetworkACLRule instead of NetworkACLRuleItem).

### Migration & mitigation
To mitigate this issue, migrate the `vpc-go-sdk` to the latest version ([v0.78.0](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.78.0)) or any of the patched versions below:
[v0.77](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.77.0)
[v0.76](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.76.5)
[v0.75](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.75.1)
[v0.74](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.74.2)
[v0.73](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.73.1)
[v0.72](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.72.1)
[v0.71](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.71.2)
[v0.70](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.70.2)
[v0.69](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.69.2)
[v0.68](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.68.1)