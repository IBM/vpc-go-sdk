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

var _ = Describe(`InstanceTemplate Unmarshalling Tests`, func() {
	Describe(`UnmarshalInstanceTemplate - Source Snapshot Detection Logic`, func() {
		Context(`When boot_volume_attachment contains source_snapshot`, func() {
			It(`Should detect source_snapshot and unmarshal as InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext`, func() {
				// Construct JSON with source_snapshot in boot_volume_attachment
				sourceSnapshotJSON := `{
					"availability_policy": {"host_failure": "restart"},
					"boot_volume_attachment": {
						"volume": {
							"name": "my-boot-volume",
							"source_snapshot": {
								"id": "0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
								"crn": "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::snapshot:0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2"
							}
						}
					},
					"confidential_compute_mode": "disabled",
					"created_at": "2024-01-01T10:00:00Z",
					"crn": "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::instance-template:0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
					"enable_secure_boot": true,
					"href": "https://us-south.iaas.cloud.ibm.com/v1/instance/templates/0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
					"id": "0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
					"name": "my-instance-template-source-snapshot",
					"profile": {"name": "bx2-2x8"},
					"resource_group": {"id": "r006-1a336ee0-0bad-4eb6-9f1c-8d3b8c5e53ed"},
					"vpc": {"id": "r006-1a336ee0-0bad-4eb6-9f1c-8d3b8c5e53ed"},
					"zone": {"name": "us-south-1"}
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(sourceSnapshotJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// This test will FAIL if the source_snapshot detection logic is missing
				// because it would return regular InstanceTemplate instead
				sourceSnapshotTemplate, isSourceSnapshotType := result.(*vpcv1.InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext)
				Expect(isSourceSnapshotType).To(BeTrue(),
					"Should be InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext type when source_snapshot is present")

				// Verify specific fields to ensure proper unmarshalling
				Expect(sourceSnapshotTemplate.ID).ToNot(BeNil())
				Expect(*sourceSnapshotTemplate.ID).To(Equal("0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2"))
				Expect(sourceSnapshotTemplate.Name).ToNot(BeNil())
				Expect(*sourceSnapshotTemplate.Name).To(Equal("my-instance-template-source-snapshot"))
				Expect(sourceSnapshotTemplate.BootVolumeAttachment).ToNot(BeNil())
			})

			It(`Should handle nested source_snapshot with minimal data`, func() {
				// Minimal JSON with just source_snapshot to test the detection logic
				minimalSourceSnapshotJSON := `{
					"boot_volume_attachment": {
						"volume": {
							"source_snapshot": {}
						}
					},
					"id": "test-id",
					"name": "test-name"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(minimalSourceSnapshotJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should detect source_snapshot even with empty object
				_, isSourceSnapshotType := result.(*vpcv1.InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext)
				Expect(isSourceSnapshotType).To(BeTrue(),
					"Should detect source_snapshot even when it's an empty object")
			})
		})

		Context(`When boot_volume_attachment does NOT contain source_snapshot`, func() {
			It(`Should unmarshal as regular InstanceTemplate`, func() {
				// Construct JSON without source_snapshot
				regularJSON := `{
					"availability_policy": {"host_failure": "restart"},
					"boot_volume_attachment": {
						"volume": {
							"name": "my-boot-volume",
							"profile": {"name": "general-purpose"}
						}
					},
					"confidential_compute_mode": "disabled",
					"created_at": "2024-01-01T10:00:00Z",
					"crn": "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::instance-template:0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
					"enable_secure_boot": true,
					"href": "https://us-south.iaas.cloud.ibm.com/v1/instance/templates/0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
					"id": "0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2",
					"name": "my-instance-template-regular",
					"profile": {"name": "bx2-2x8"},
					"resource_group": {"id": "r006-1a336ee0-0bad-4eb6-9f1c-8d3b8c5e53ed"},
					"vpc": {"id": "r006-1a336ee0-0bad-4eb6-9f1c-8d3b8c5e53ed"},
					"zone": {"name": "us-south-1"}
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(regularJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify it's the regular InstanceTemplate type
				regularTemplate, isRegularType := result.(*vpcv1.InstanceTemplate)
				Expect(isRegularType).To(BeTrue(),
					"Should be InstanceTemplate type when source_snapshot is not present")

				// Verify specific fields
				Expect(regularTemplate.ID).ToNot(BeNil())
				Expect(*regularTemplate.ID).To(Equal("0717-e6c2c7d8-ad57-4f38-a21c-a86265b6aeb2"))
				Expect(regularTemplate.Name).ToNot(BeNil())
				Expect(*regularTemplate.Name).To(Equal("my-instance-template-regular"))
			})
		})

		Context(`Edge cases for source_snapshot detection`, func() {
			It(`Should handle null boot_volume_attachment`, func() {
				nullBootVolumeJSON := `{
					"boot_volume_attachment": null,
					"id": "test-id",
					"name": "test-name"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(nullBootVolumeJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should be regular InstanceTemplate when boot_volume_attachment is null
				_, isRegularType := result.(*vpcv1.InstanceTemplate)
				Expect(isRegularType).To(BeTrue(),
					"Should be InstanceTemplate type when boot_volume_attachment is null")
			})

			It(`Should handle missing boot_volume_attachment field`, func() {
				noBootVolumeJSON := `{
					"id": "test-id",
					"name": "test-name"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(noBootVolumeJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should be regular InstanceTemplate when boot_volume_attachment is missing
				_, isRegularType := result.(*vpcv1.InstanceTemplate)
				Expect(isRegularType).To(BeTrue(),
					"Should be InstanceTemplate type when boot_volume_attachment is missing")
			})

			It(`Should handle empty boot_volume_attachment object`, func() {
				emptyBootVolumeJSON := `{
					"boot_volume_attachment": {},
					"id": "test-id",
					"name": "test-name"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(emptyBootVolumeJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should be regular InstanceTemplate when boot_volume_attachment is empty
				_, isRegularType := result.(*vpcv1.InstanceTemplate)
				Expect(isRegularType).To(BeTrue(),
					"Should be InstanceTemplate type when boot_volume_attachment is empty")
			})

			It(`Should handle boot_volume_attachment with null volume`, func() {
				nullVolumeJSON := `{
					"boot_volume_attachment": {
						"volume": null
					},
					"id": "test-id",
					"name": "test-name"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(nullVolumeJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should be regular InstanceTemplate when volume is null
				_, isRegularType := result.(*vpcv1.InstanceTemplate)
				Expect(isRegularType).To(BeTrue(),
					"Should be InstanceTemplate type when volume is null")
			})

			It(`Should handle boot_volume_attachment with empty volume object`, func() {
				emptyVolumeJSON := `{
					"boot_volume_attachment": {
						"volume": {}
					},
					"id": "test-id",
					"name": "test-name"
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(emptyVolumeJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should be regular InstanceTemplate when volume is empty (no source_snapshot field)
				_, isRegularType := result.(*vpcv1.InstanceTemplate)
				Expect(isRegularType).To(BeTrue(),
					"Should be InstanceTemplate type when volume is empty (no source_snapshot)")
			})
		})

		Context(`Comprehensive field unmarshalling tests`, func() {
			It(`Should unmarshal all fields correctly for source_snapshot type`, func() {
				comprehensiveSourceSnapshotJSON := `{
					"availability_policy": {"host_failure": "restart"},
					"boot_volume_attachment": {
						"volume": {
							"name": "boot-vol",
							"source_snapshot": {"id": "snapshot-123"}
						}
					},
					"cluster_network_attachments": [{"id": "cluster-net-1"}],
					"confidential_compute_mode": "disabled",
					"created_at": "2024-01-01T10:00:00Z",
					"crn": "crn:test",
					"default_trusted_profile": {"auto_link": false},
					"enable_secure_boot": true,
					"href": "https://test.com",
					"id": "template-123",
					"keys": [{"id": "key-1"}],
					"metadata_service": {"enabled": true},
					"name": "comprehensive-template",
					"placement_target": {"id": "placement-1"},
					"profile": {"name": "test-profile"},
					"reservation_affinity": {"policy": "manual"},
					"resource_group": {"id": "rg-1"},
					"total_volume_bandwidth": 500,
					"user_data": "test-user-data",
					"volume_attachments": [{"id": "vol-attach-1"}],
					"volume_bandwidth_qos_mode": "pooled",
					"vpc": {"id": "vpc-1"},
					"network_attachments": [{"id": "net-attach-1"}],
					"network_interfaces": [{"id": "nic-1"}],
					"primary_network_attachment": {"id": "primary-net-attach"},
					"zone": {"name": "us-south-1"},
					"primary_network_interface": {"id": "primary-nic"}
				}`

				var rawMap map[string]json.RawMessage
				err := json.Unmarshal([]byte(comprehensiveSourceSnapshotJSON), &rawMap)
				Expect(err).To(BeNil())

				var result interface{}
				err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())

				// Should be source snapshot type
				sourceSnapshotTemplate, isSourceSnapshotType := result.(*vpcv1.InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext)
				Expect(isSourceSnapshotType).To(BeTrue())

				// Verify various fields are properly unmarshalled
				Expect(sourceSnapshotTemplate.ID).ToNot(BeNil())
				Expect(*sourceSnapshotTemplate.ID).To(Equal("template-123"))
				Expect(sourceSnapshotTemplate.Name).ToNot(BeNil())
				Expect(*sourceSnapshotTemplate.Name).To(Equal("comprehensive-template"))
				Expect(sourceSnapshotTemplate.ConfidentialComputeMode).ToNot(BeNil())
				Expect(*sourceSnapshotTemplate.ConfidentialComputeMode).To(Equal("disabled"))
				Expect(sourceSnapshotTemplate.EnableSecureBoot).ToNot(BeNil())
				Expect(*sourceSnapshotTemplate.EnableSecureBoot).To(BeTrue())
				Expect(sourceSnapshotTemplate.BootVolumeAttachment).ToNot(BeNil())
			})
		})
	})

	Describe(`Integration Test - This test will FAIL if source_snapshot detection logic is missing`, func() {
		It(`Should demonstrate that the test detects missing source_snapshot logic`, func() {
			// This is the imp test that validates the source_snapshot detection logic
			// If the logic is removed from UnmarshalInstanceTemplate, this test will fail
			sourceSnapshotJSON := `{
				"boot_volume_attachment": {
					"volume": {
						"source_snapshot": {"id": "test-snapshot"}
					}
				},
				"id": "test-template",
				"name": "test-name"
			}`

			var rawMap map[string]json.RawMessage
			err := json.Unmarshal([]byte(sourceSnapshotJSON), &rawMap)
			Expect(err).To(BeNil())

			var result interface{}
			err = vpcv1.UnmarshalInstanceTemplate(rawMap, &result)
			Expect(err).To(BeNil())

			// THE IMP ASSERTION: This will fail if source_snapshot detection is missing
			_, isSourceSnapshotType := result.(*vpcv1.InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext)
			Expect(isSourceSnapshotType).To(BeTrue(),
				"CRITICAL: source_snapshot detection logic appears to be missing! "+
					"Expected InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext but got regular InstanceTemplate")
		})
	})
})
