# Known Issues for vpc-go-sdk

## Security Group Rules and Network ACL Rules - discriminator issue
#### Publication date - "2025-12-18"

### Patch fix versions available in the following versions - 
[v0.76.5](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.76.5)
[v0.75.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.75.1)
[v0.74.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.74.2)
[v0.73.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.73.1)
[v0.72.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.72.1)
[v0.71.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.71.2)
[v0.70.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.70.2)
[v0.69.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.69.2)
[v0.68.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.68.1)
and later versions

### Affected component
vpc-go-sdk

### Affected operations
Security group rules and Network ACL rules 

### Issue Summary
#### Support for `any` Protocol in Security Group Rules and Network ACL Rules
Security Group rules and Network ACL rules now support the new protocol value `any`. Earlier versions of SDK could result in following error reading a rule with an unsupported protocol value. 
```
error unmarshalling vpcv1.SecurityGroupCollection: error unmarshalling property 'security_groups' as []vpcv1.SecurityGroup: error unmarshalling property 'rules' as []vpcv1.SecurityGroupRuleIntf: unrecognized value for discriminator property 'protocol': any
```

Fix aligns fallback behavior and error identifiers with the correct model name ( e.g. NetworkACLRule instead of NetworkACLRuleItem).

### Migration & mitigation
Migrate to patched versions or later. [v0.76.5](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.76.5)
[v0.75.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.75.1)
[v0.74.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.74.2)
[v0.73.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.73.1)
[v0.72.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.72.1)
[v0.71.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.71.2)
[v0.70.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.70.2)
[v0.69.2](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.69.2)
[v0.68.1](https://github.com/IBM/vpc-go-sdk/releases/tag/v0.68.1)

