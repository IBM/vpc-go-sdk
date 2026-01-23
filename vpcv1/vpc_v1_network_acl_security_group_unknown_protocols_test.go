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

var _ = Describe(`NetworkACL and SecurityGroup Rule Unmarshalling Tests - Protocol 'all' Deprecation [protocol-fallback-new]`, func() {
	Describe(`UnmarshalNetworkACLRuleItem - 'all' Protocol Deprecation`, func() {
		Context(`When 'all' protocol is deprecated and should use generic fallback`, func() {
			It(`Should unmarshal 'all' protocol to generic type (no longer specific type)`, func() {
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

				// CRITICAL: 'all' should now route to generic, not specific type
				genericRule, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(BeTrue(), "'all' protocol should now use generic NetworkACLRuleItem type")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("all"))
				Expect(genericRule.Action).ToNot(BeNil())
				Expect(*genericRule.Action).To(Equal("allow"))
			})

			It(`Should handle 'all' protocol with all base fields populated`, func() {
				allProtocolWithFieldsJSON := `{
					"action": "deny",
					"before": {
						"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/before-rule",
						"id": "before-rule-id"
					},
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "192.168.0.0/16",
					"direction": "outbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule2",
					"id": "rule-id-2",
					"ip_version": "ipv4",
					"name": "deny-all-traffic",
					"protocol": "all",
					"source": "10.0.0.0/8"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(allProtocolWithFieldsJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(BeTrue())
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("all"))
				Expect(genericRule.Action).ToNot(BeNil())
				Expect(*genericRule.Action).To(Equal("deny"))
				Expect(genericRule.Destination).ToNot(BeNil())
				Expect(*genericRule.Destination).To(Equal("192.168.0.0/16"))
			})
		})

		Context(`Legacy protocols (icmp, tcp, udp) continue to use their specific types`, func() {
			It(`Should still unmarshal 'icmp' to its specific type`, func() {
				icmpProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule3",
					"id": "rule-id-3",
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

				// ICMP should still use specific type
				icmpRule, isIcmpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp)
				Expect(isIcmpType).To(BeTrue(), "ICMP should still use specific type")
				Expect(icmpRule.Protocol).ToNot(BeNil())
				Expect(*icmpRule.Protocol).To(Equal("icmp"))
			})

			It(`Should still unmarshal 'tcp' to its specific type`, func() {
				tcpProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule4",
					"id": "rule-id-4",
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

				// TCP should still use specific type
				tcpRule, isTcpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
				Expect(isTcpType).To(BeTrue(), "TCP should still use specific type")
				Expect(tcpRule.Protocol).ToNot(BeNil())
				Expect(*tcpRule.Protocol).To(Equal("tcp"))
			})

			It(`Should still unmarshal 'udp' to its specific type`, func() {
				udpProtocolJSON := `{
					"action": "allow",
					"created_at": "2024-01-15T10:30:00Z",
					"destination": "0.0.0.0/0",
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/rule5",
					"id": "rule-id-5",
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

				// UDP should still use specific type
				udpRule, isUdpType := result.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
				Expect(isUdpType).To(BeTrue(), "UDP should still use specific type")
				Expect(udpRule.Protocol).ToNot(BeNil())
				Expect(*udpRule.Protocol).To(Equal("udp"))
			})
		})
	})

	Describe(`UnmarshalNetworkACLRule - 'all' Protocol Deprecation`, func() {
		Context(`When 'all' protocol should use generic fallback`, func() {
			It(`Should unmarshal 'all' protocol to generic type`, func() {
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

				var result vpcv1.NetworkACLRuleIntf
				err = vpcv1.UnmarshalNetworkACLRule(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.NetworkACLRule)
				Expect(isGenericType).To(BeTrue(), "'all' should use generic type")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("all"))
			})
		})

		Context(`Legacy protocols continue to use specific types`, func() {
			It(`Should still unmarshal 'tcp' to its specific type`, func() {
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

				var result vpcv1.NetworkACLRuleIntf
				err = vpcv1.UnmarshalNetworkACLRule(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				tcpRule, isTcpType := result.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolTcpudp)
				Expect(isTcpType).To(BeTrue())
				Expect(tcpRule.Protocol).ToNot(BeNil())
				Expect(*tcpRule.Protocol).To(Equal("tcp"))
			})
		})
	})

	Describe(`UnmarshalSecurityGroupRule - 'all' Protocol Deprecation`, func() {
		Context(`When 'all' protocol should use generic fallback`, func() {
			It(`Should unmarshal 'all' protocol to generic type`, func() {
				allProtocolJSON := `{
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/sg1/rules/rule1",
					"id": "sg-rule-id-1",
					"ip_version": "ipv4",
					"protocol": "all",
					"remote": {
						"cidr_block": "0.0.0.0/0"
					}
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(allProtocolJSON), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.SecurityGroupRuleIntf
				err = vpcv1.UnmarshalSecurityGroupRule(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				genericRule, isGenericType := result.(*vpcv1.SecurityGroupRule)
				Expect(isGenericType).To(BeTrue(), "'all' should use generic type")
				Expect(genericRule.Protocol).ToNot(BeNil())
				Expect(*genericRule.Protocol).To(Equal("all"))
			})
		})

		Context(`Legacy protocols continue to use specific types`, func() {
			It(`Should still unmarshal 'tcp' to its specific type`, func() {
				tcpProtocolJSON := `{
					"direction": "inbound",
					"href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/sg2/rules/rule2",
					"id": "sg-rule-id-2",
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
				Expect(isTcpType).To(BeTrue())
				Expect(tcpRule.Protocol).ToNot(BeNil())
				Expect(*tcpRule.Protocol).To(Equal("tcp"))
			})
		})
	})

	Describe(`Integration Test - 'all' Protocol Migration [protocol-fallback-new]`, func() {
		It(`Should demonstrate ONLY 'all' protocol routes to generic type`, func() {
			// CRITICAL: This test validates that ONLY 'all' is routed to generic
			allJSON := `{"action": "allow", "protocol": "all", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8"}`
			var rawMap map[string]json.RawMessage
			err := json.Unmarshal([]byte(allJSON), &rawMap)
			Expect(err).To(BeNil())

			var result vpcv1.NetworkACLRuleItemIntf
			err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
			Expect(err).To(BeNil())

			_, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
			Expect(isGenericType).To(BeTrue(), "CRITICAL: ONLY 'all' protocol should use generic type")
		})

		It(`Should verify icmp/tcp/udp still use specific types (NOT generic)`, func() {
			// Test that icmp, tcp, udp are NOT routed to generic
			protocols := map[string]string{
				"icmp": `{"action": "allow", "protocol": "icmp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8", "code": 0, "type": 8}`,
				"tcp":  `{"action": "allow", "protocol": "tcp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8", "destination_port_min": 443, "destination_port_max": 443}`,
				"udp":  `{"action": "allow", "protocol": "udp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8", "destination_port_min": 53, "destination_port_max": 53}`,
			}

			for protocol, jsonStr := range protocols {
				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(jsonStr), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())

				// These should NOT be generic type
				_, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(BeFalse(), "Protocol %s should NOT use generic type", protocol)
			}
		})

		It(`Should maintain complete backward compatibility except for 'all'`, func() {
			// Comprehensive test showing the migration impact
			testCases := []struct {
				protocol           string
				json               string
				shouldBeGeneric    bool
				expectedTypeString string
			}{
				{
					protocol:           "all",
					json:               `{"action": "allow", "protocol": "all", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8"}`,
					shouldBeGeneric:    true,
					expectedTypeString: "*vpcv1.NetworkACLRuleItem",
				},
				{
					protocol:           "icmp",
					json:               `{"action": "allow", "protocol": "icmp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8", "code": 0, "type": 8}`,
					shouldBeGeneric:    false,
					expectedTypeString: "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp",
				},
				{
					protocol:           "tcp",
					json:               `{"action": "allow", "protocol": "tcp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8", "destination_port_min": 443, "destination_port_max": 443}`,
					shouldBeGeneric:    false,
					expectedTypeString: "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp",
				},
				{
					protocol:           "udp",
					json:               `{"action": "allow", "protocol": "udp", "destination": "0.0.0.0/0", "direction": "inbound", "id": "test", "source": "10.0.0.0/8", "destination_port_min": 53, "destination_port_max": 53}`,
					shouldBeGeneric:    false,
					expectedTypeString: "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp",
				},
			}

			for _, tc := range testCases {
				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(tc.json), &rawMap)
				Expect(err).To(BeNil())

				var result vpcv1.NetworkACLRuleItemIntf
				err = vpcv1.UnmarshalNetworkACLRuleItem(rawMap, &result)
				Expect(err).To(BeNil())

				_, isGenericType := result.(*vpcv1.NetworkACLRuleItem)
				Expect(isGenericType).To(Equal(tc.shouldBeGeneric),
					"Protocol %s generic type expectation mismatch", tc.protocol)
			}
		})
	})
})
