/**
 * (C) Copyright IBM Corp. 2023, 2024, 2025.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package vpcv1_test

import (
	"encoding/json"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`NetworkACL and SecurityGroup Rule Unmarshalling Tests - Unknown Protocol Support [protocol-fallback]`, func() {
	Describe(`UnmarshalNetworkACLRuleItem - Unknown Protocol Fallback`, func() {
		Context(`When protocol is a known value (all, icmp, tcp, udp)`, func() {
			It(`Should unmarshal 'all' protocol to specific type`, func() {
				allProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule1",
					"id": "rule-id-1",
					"ip_version": "ipv4",
					"name": "allow-all",
					"protocol": "all",
					"source": "10.0.0.0/8"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(allProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify it's the correct protocol-specific type
				_, isProtocolAll := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAll)
				Expect(isProtocolAll).To(BeTrue(), "Should be NetworkACLRuleProtocolAll type")
			})

			It(`Should unmarshal 'tcp' protocol to specific type`, func() {
				tcpProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule2",
					"id": "rule-id-2",
					"ip_version": "ipv4",
					"name": "allow-tcp",
					"protocol": "tcp",
					"source": "10.0.0.0/8",
					"destination_port_min": 443,
					"destination_port_max": 443
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(tcpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify it's the correct protocol-specific type
				tcpRule, isTcpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
				Expect(isTcpType).To(BeTrue(), "Should be NetworkACLRuleProtocolTcpudp type")
				Expect(tcpRule.Protocol).ToNot(BeNil())
				Expect(*tcpRule.Protocol).To(Equal("tcp"))
			})

			It(`Should unmarshal 'udp' protocol to specific type`, func() {
				udpProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule3",
					"id": "rule-id-3",
					"ip_version": "ipv4",
					"name": "allow-udp",
					"protocol": "udp",
					"source": "10.0.0.0/8",
					"destination_port_min": 53,
					"destination_port_max": 53
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(udpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify it's the correct protocol-specific type
				udpRule, isUdpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
				Expect(isUdpType).To(BeTrue(), "Should be NetworkACLRuleProtocolTcpudp type")
				Expect(udpRule.Protocol).ToNot(BeNil())
				Expect(*udpRule.Protocol).To(Equal("udp"))
			})

			It(`Should unmarshal 'icmp' protocol to specific type`, func() {
				icmpProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule4",
					"id": "rule-id-4",
					"ip_version": "ipv4",
					"name": "allow-icmp",
					"protocol": "icmp",
					"source": "10.0.0.0/8",
					"code": 0,
					"type": 8
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(icmpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify it's the correct protocol-specific type
				icmpRule, isIcmpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp)
				Expect(isIcmpType).To(BeTrue(), "Should be NetworkACLRuleProtocolIcmp type")
				Expect(icmpRule.Protocol).ToNot(BeNil())
				Expect(*icmpRule.Protocol).To(Equal("icmp"))
			})
		})

		Context(`When protocol is an unknown value (any, esp, sctp)`, func() {
			It(`Should gracefully handle 'any' protocol without error`, func() {
				anyProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule5",
					"id": "rule-id-5",
					"ip_version": "ipv4",
					"name": "allow-any-protocol",
					"protocol": "any",
					"source": "10.0.0.0/8"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(anyProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)

				// CRITICAL: Should not error for unknown protocols
				Expect(err).To(BeNil(), "Should not error for 'any' protocol")
				Expect(result).ToNot(BeNil())

				// Verify it falls back to generic NetworkACLRuleItem
				genericRule, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(BeTrue(), "Should fallback to generic NetworkACLRuleItem for unknown protocol")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("any"))
				Expect(genericRule.Action).ToNot(BeNil())
				Expect(*genericRule.Action).To(Equal("allow"))
			})

			It(`Should gracefully handle 'esp' protocol without error`, func() {
				espProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "192.168.1.0/24",
					"direction": "outbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule6",
					"id": "rule-id-6",
					"ip_version": "ipv4",
					"name": "allow-esp-vpn",
					"protocol": "esp",
					"source": "10.0.0.0/16"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(espProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)

				// CRITICAL: Should not error for unknown protocols
				Expect(err).To(BeNil(), "Should not error for 'esp' protocol")
				Expect(result).ToNot(BeNil())

				// Verify it falls back to generic NetworkACLRuleItem
				genericRule, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(BeTrue(), "Should fallback to generic NetworkACLRuleItem for unknown protocol")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("esp"))
				Expect(genericRule.Direction).ToNot(BeNil())
				Expect(*genericRule.Direction).To(Equal("outbound"))
			})

			It(`Should gracefully handle 'sctp' protocol without error`, func() {
				sctpProtocolJSON := `{
					"action": "deny",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "172.16.0.0/12",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule7",
					"id": "rule-id-7",
					"ip_version": "ipv4",
					"name": "deny-sctp",
					"protocol": "sctp",
					"source": "0.0.0.0/0",
					"destination_port_min": 1024,
					"destination_port_max": 65535
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(sctpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)

				// CRITICAL: Should not error for unknown protocols
				Expect(err).To(BeNil(), "Should not error for 'sctp' protocol")
				Expect(result).ToNot(BeNil())

				// Verify it falls back to generic NetworkACLRuleItem
				genericRule, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(BeTrue(), "Should fallback to generic NetworkACLRuleItem for unknown protocol")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("sctp"))

				// Verify optional port fields are captured
				Expect(genericRule.DestinationPortMin).ToNot(BeNil())
				Expect(*genericRule.DestinationPortMin).To(Equal(int64(1024)))
				Expect(genericRule.DestinationPortMax).ToNot(BeNil())
				Expect(*genericRule.DestinationPortMax).To(Equal(int64(65535)))
			})
		})
	})

	Describe(`UnmarshalNetworkACLRule - Unknown Protocol Fallback`, func() {
		Context(`When protocol is an unknown value`, func() {
			It(`Should gracefully handle 'any' protocol`, func() {
				anyProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule1",
					"id": "rule-id-1",
					"ip_version": "ipv4",
					"name": "allow-any",
					"protocol": "any",
					"source": "10.0.0.0/8"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(anyProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleIntf
				err = vpcv1.UnmarshalNetworkACLRule(rawMap, &result)

				Expect(err).To(BeNil(), "Should not error for 'any' protocol")
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.NetworkACLRule)
				Expect(isGenericType).To(BeTrue())
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("any"))
			})

			It(`Should gracefully handle 'esp' protocol`, func() {
				espProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "192.168.1.0/24",
					"direction": "outbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule2",
					"id": "rule-id-2",
					"ip_version": "ipv4",
					"name": "allow-esp",
					"protocol": "esp",
					"source": "10.0.0.0/16"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(espProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleIntf
				err = vpcv1.UnmarshalNetworkACLRule(rawMap, &result)

				Expect(err).To(BeNil(), "Should not error for 'esp' protocol")
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.NetworkACLRule)
				Expect(isGenericType).To(BeTrue())
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("esp"))
			})

			It(`Should gracefully handle 'sctp' protocol with port fields`, func() {
				sctpProtocolJSON := `{
					"action": "deny",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "172.16.0.0/12",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule3",
					"id": "rule-id-3",
					"ip_version": "ipv4",
					"name": "deny-sctp",
					"protocol": "sctp",
					"source": "0.0.0.0/0",
					"destination_port_min": 2000,
					"destination_port_max": 3000
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(sctpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleIntf
				err = vpcv1.UnmarshalNetworkACLRule(rawMap, &result)

				Expect(err).To(BeNil(), "Should not error for 'sctp' protocol")
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.NetworkACLRule)
				Expect(isGenericType).To(BeTrue())
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("sctp"))
				Expect(genericRule.DestinationPortMin).ToNot(BeNil())
				Expect(*genericRule.DestinationPortMin).To(Equal(int64(2000)))
			})
		})
	})

	Describe(`UnmarshalSecurityGroupRule - Unknown Protocol Fallback`, func() {
		Context(`When protocol is a known value`, func() {
			It(`Should unmarshal 'tcp' protocol to specific type`, func() {
				tcpProtocolJSON := `{
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/sg1/rules/rule1",
					"id": "sg-rule-id-1",
					"ip_version": "ipv4",
					"protocol": "tcp",
					"remote": {
						"cidr_block": "0.0.0.0/0"
					},
					"port_min": 443,
					"port_max": 443
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(tcpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.SecurityGroupRuleIntf
				err = vpcv1.UnmarshalSecurityGroupRule(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				tcpRule, isTcpType := result.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
				Expect(isTcpType).To(BeTrue(), "Should be SecurityGroupRuleProtocolTcpudp type")
				Expect(tcpRule.Protocol).ToNot(BeNil())
				Expect(*tcpRule.Protocol).To(Equal("tcp"))
			})
		})

		Context(`When protocol is an unknown value`, func() {
			It(`Should gracefully handle 'any' protocol`, func() {
				anyProtocolJSON := `{
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/sg1/rules/rule1",
					"id": "sg-rule-id-1",
					"ip_version": "ipv4",
					"protocol": "any",
					"remote": {
						"cidr_block": "0.0.0.0/0"
					}
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(anyProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.SecurityGroupRuleIntf
				err = vpcv1.UnmarshalSecurityGroupRule(rawMap, &result)

				Expect(err).To(BeNil(), "Should not error for 'any' protocol")
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.SecurityGroupRule)
				Expect(isGenericType).To(BeTrue(), "Should fallback to generic SecurityGroupRule")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("any"))
			})

			It(`Should gracefully handle 'esp' protocol`, func() {
				espProtocolJSON := `{
					"direction": "outbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/sg2/rules/rule2",
					"id": "sg-rule-id-2",
					"ip_version": "ipv4",
					"protocol": "esp",
					"remote": {
						"cidr_block": "192.168.0.0/16"
					}
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(espProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.SecurityGroupRuleIntf
				err = vpcv1.UnmarshalSecurityGroupRule(rawMap, &result)

				Expect(err).To(BeNil(), "Should not error for 'esp' protocol")
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.SecurityGroupRule)
				Expect(isGenericType).To(BeTrue())
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("esp"))
			})

			It(`Should gracefully handle 'sctp' protocol with port fields`, func() {
				sctpProtocolJSON := `{
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/sg3/rules/rule3",
					"id": "sg-rule-id-3",
					"ip_version": "ipv4",
					"protocol": "sctp",
					"remote": {
						"cidr_block": "10.0.0.0/8"
					},
					"port_min": 3868,
					"port_max": 3868
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(sctpProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.SecurityGroupRuleIntf
				err = vpcv1.UnmarshalSecurityGroupRule(rawMap, &result)

				Expect(err).To(BeNil(), "Should not error for 'sctp' protocol")
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.SecurityGroupRule)
				Expect(isGenericType).To(BeTrue())
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("sctp"))

				// Verify port fields are captured
				Expect(genericRule.PortMin).ToNot(BeNil())
				Expect(*genericRule.PortMin).To(Equal(int64(3868)))
				Expect(genericRule.PortMax).ToNot(BeNil())
				Expect(*genericRule.PortMax).To(Equal(int64(3868)))
			})
		})
	})

	Describe(`Integration Test - Verifying Unknown Protocol Support`, func() {
		It(`Should demonstrate backward compatibility - known protocols still work`, func() {
			// Test that existing functionality isn't broken
			tcpJSON := `{"action": "allow", "protocol": "tcp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8"}`
			var rawMap map[string]json.RawMessage
			err := json.Unmarshal([]byte(tcpJSON), &rawMap)
			Expect(err).To(BeNil())

			var result vpcv1.NetworkACLRuleItemIntf
			err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
			Expect(err).To(BeNil())

			_, isTcpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
			Expect(isTcpType).To(BeTrue(), "Known protocols should still use specific types")
		})

		It(`Should demonstrate new capability - unknown protocols handled gracefully`, func() {
			// CRITICAL TEST: This validates the unknown protocol fallback logic
			// Without the generic unmarshal functions, this would error out
			unknownJSON := `{"action": "allow", "protocol": "gre", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8"}`
			var rawMap map[string]json.RawMessage
			err := json.Unmarshal([]byte(unknownJSON), &rawMap)
			Expect(err).To(BeNil())

			var result vpcv1.NetworkACLRuleItemIntf
			err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)

			Expect(err).To(BeNil(), "CRITICAL: Unknown protocol 'gre' should not cause error")

			genericRule, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
			Expect(isGenericType).To(BeTrue(), "Unknown protocols should fallback to generic type")
			Expect(genericRule.Protocol).ToNot(BeNil())
			Expect(*genericRule.Protocol).To(Equal("gre"))
		})
	})
})
