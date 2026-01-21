package vpcv1_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestInstancePrototypeMarshalBug(t *testing.T) {
	// Create the prototype exactly like Karpenter does
	imageID := "r010-test-image-id"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"
	subnetID := "test-subnet-id"

	// Create VNI prototype
	vniPrototype := &vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContext{
		Subnet: &vpcv1.SubnetIdentityByID{
			ID: &subnetID,
		},
	}

	primaryNetworkAttachment := &vpcv1.InstanceNetworkAttachmentPrototype{
		VirtualNetworkInterface: vniPrototype,
	}

	// Create instance prototype
	instancePrototype := &vpcv1.InstancePrototypeInstanceByImageInstanceByImageInstanceByNetworkAttachment{
		Image:                    &vpcv1.ImageIdentityByID{ID: &imageID},
		Zone:                     &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile:                  &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:                      &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:                     &instanceName,
		PrimaryNetworkAttachment: primaryNetworkAttachment,
	}

	// Verify fields are set
	if instancePrototype.Image == nil {
		t.Fatal("Image field should not be nil")
	}
	if instancePrototype.Zone == nil {
		t.Fatal("Zone field should not be nil")
	}
	if instancePrototype.VPC == nil {
		t.Fatal("VPC field should not be nil")
	}
	if instancePrototype.Profile == nil {
		t.Fatal("Profile field should not be nil")
	}

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(instancePrototype, "", "  ")
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	t.Logf("Marshaled JSON:\n%s\n", jsonString)

	// These assertions will FAIL if the bug exists
	if !strings.Contains(jsonString, `"image"`) {
		t.Error("JSON should contain image field - BUG CONFIRMED")
	}
	if !strings.Contains(jsonString, `"zone"`) {
		t.Error("JSON should contain zone field - BUG CONFIRMED")
	}
	if !strings.Contains(jsonString, `"vpc"`) {
		t.Error("JSON should contain vpc field - BUG CONFIRMED")
	}
	if !strings.Contains(jsonString, `"profile"`) {
		t.Error("JSON should contain profile field - BUG CONFIRMED")
	}

	// Verify actual values
	if !strings.Contains(jsonString, imageID) {
		t.Error("JSON should contain image ID value")
	}
	if !strings.Contains(jsonString, zoneName) {
		t.Error("JSON should contain zone name value")
	}
	if !strings.Contains(jsonString, vpcID) {
		t.Error("JSON should contain VPC ID value")
	}
	if !strings.Contains(jsonString, profileName) {
		t.Error("JSON should contain profile name value")
	}
}

func TestInstancePrototypeInterfaceMarshalBug(t *testing.T) {
	// This test verifies that marshaling through the InstancePrototypeIntf interface
	// correctly calls the concrete type's MarshalJSON method.
	// This is the CRITICAL fix - without it, interface marshaling drops fields.

	imageID := "r010-test-image-id"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"
	subnetID := "test-subnet-id"

	// Create VNI prototype
	vniPrototype := &vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContext{
		Subnet: &vpcv1.SubnetIdentityByID{
			ID: &subnetID,
		},
	}

	primaryNetworkAttachment := &vpcv1.InstanceNetworkAttachmentPrototype{
		VirtualNetworkInterface: vniPrototype,
	}

	// Create concrete prototype
	concretePrototype := &vpcv1.InstancePrototypeInstanceByImageInstanceByImageInstanceByNetworkAttachment{
		Image:                    &vpcv1.ImageIdentityByID{ID: &imageID},
		Zone:                     &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile:                  &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:                      &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:                     &instanceName,
		PrimaryNetworkAttachment: primaryNetworkAttachment,
	}

	// Cast to interface type (THIS is how the SDK uses it)
	var interfacePrototype vpcv1.InstancePrototypeIntf = concretePrototype

	// Marshal through the INTERFACE (critical test)
	jsonBytes, err := json.MarshalIndent(interfacePrototype, "", "  ")
	if err != nil {
		t.Fatalf("JSON marshaling through interface failed: %v", err)
	}

	jsonString := string(jsonBytes)
	t.Logf("Marshaled JSON through interface:\n%s\n", jsonString)

	// These assertions verify the interface marshaling works
	if !strings.Contains(jsonString, `"image"`) {
		t.Error("JSON should contain image field when marshaled through interface - INTERFACE BUG CONFIRMED")
	}
	if !strings.Contains(jsonString, `"zone"`) {
		t.Error("JSON should contain zone field when marshaled through interface - INTERFACE BUG CONFIRMED")
	}
	if !strings.Contains(jsonString, `"vpc"`) {
		t.Error("JSON should contain vpc field when marshaled through interface - INTERFACE BUG CONFIRMED")
	}
	if !strings.Contains(jsonString, `"profile"`) {
		t.Error("JSON should contain profile field when marshaled through interface - INTERFACE BUG CONFIRMED")
	}

	// Verify actual values
	if !strings.Contains(jsonString, imageID) {
		t.Error("JSON should contain image ID value when marshaled through interface")
	}
	if !strings.Contains(jsonString, zoneName) {
		t.Error("JSON should contain zone name value when marshaled through interface")
	}
	if !strings.Contains(jsonString, vpcID) {
		t.Error("JSON should contain VPC ID value when marshaled through interface")
	}
	if !strings.Contains(jsonString, profileName) {
		t.Error("JSON should contain profile name value when marshaled through interface")
	}
}

func TestBaseInstancePrototypeMarshal(t *testing.T) {
	// Test the base InstancePrototype type marshaling
	imageID := "r010-test-image-id"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"

	basePrototype := &vpcv1.InstancePrototype{
		Image:   &vpcv1.ImageIdentityByID{ID: &imageID},
		Zone:    &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile: &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:     &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:    &instanceName,
	}

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(basePrototype, "", "  ")
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	t.Logf("Marshaled base prototype JSON:\n%s\n", jsonString)

	// Verify interface fields are present
	if !strings.Contains(jsonString, `"image"`) {
		t.Error("Base prototype JSON should contain image field")
	}
	if !strings.Contains(jsonString, `"zone"`) {
		t.Error("Base prototype JSON should contain zone field")
	}
	if !strings.Contains(jsonString, `"vpc"`) {
		t.Error("Base prototype JSON should contain vpc field")
	}
	if !strings.Contains(jsonString, `"profile"`) {
		t.Error("Base prototype JSON should contain profile field")
	}
}

func TestInstancePrototypeInstanceByImageMarshal(t *testing.T) {
	imageID := "r010-test-image-id"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"

	prototype := &vpcv1.InstancePrototypeInstanceByImage{
		Image:   &vpcv1.ImageIdentityByID{ID: &imageID},
		Zone:    &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile: &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:     &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:    &instanceName,
	}

	jsonBytes, err := json.Marshal(prototype)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	if !strings.Contains(jsonString, `"image"`) || !strings.Contains(jsonString, `"zone"`) ||
		!strings.Contains(jsonString, `"vpc"`) || !strings.Contains(jsonString, `"profile"`) {
		t.Errorf("InstancePrototypeInstanceByImage missing interface fields in JSON: %s", jsonString)
	}
}

func TestInstancePrototypeInstanceByCatalogOfferingMarshal(t *testing.T) {
	catalogOfferingVersionCRN := "crn:test:offering"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"

	prototype := &vpcv1.InstancePrototypeInstanceByCatalogOffering{
		CatalogOffering: &vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByVersion{
			Version: &vpcv1.CatalogOfferingVersionIdentityCatalogOfferingVersionByCRN{
				CRN: &catalogOfferingVersionCRN,
			},
		},
		Zone:    &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile: &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:     &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:    &instanceName,
	}

	jsonBytes, err := json.Marshal(prototype)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	if !strings.Contains(jsonString, `"catalog_offering"`) || !strings.Contains(jsonString, `"zone"`) ||
		!strings.Contains(jsonString, `"vpc"`) || !strings.Contains(jsonString, `"profile"`) {
		t.Errorf("InstancePrototypeInstanceByCatalogOffering missing interface fields in JSON: %s", jsonString)
	}
}

func TestInstancePrototypeInstanceBySourceSnapshotMarshal(t *testing.T) {
	snapshotID := "r010-test-snapshot-id"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"

	prototype := &vpcv1.InstancePrototypeInstanceBySourceSnapshot{
		BootVolumeAttachment: &vpcv1.VolumeAttachmentPrototypeInstanceBySourceSnapshotContext{
			Volume: &vpcv1.VolumePrototypeInstanceBySourceSnapshotContext{
				SourceSnapshot: &vpcv1.SnapshotIdentityByID{ID: &snapshotID},
			},
		},
		Zone:    &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile: &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:     &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:    &instanceName,
	}

	jsonBytes, err := json.Marshal(prototype)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	if !strings.Contains(jsonString, `"boot_volume_attachment"`) || !strings.Contains(jsonString, `"zone"`) ||
		!strings.Contains(jsonString, `"vpc"`) || !strings.Contains(jsonString, `"profile"`) {
		t.Errorf("InstancePrototypeInstanceBySourceSnapshot missing interface fields in JSON: %s", jsonString)
	}
}

func TestInstancePrototypeInstanceBySourceTemplateMarshal(t *testing.T) {
	templateID := "r010-test-template-id"
	zoneName := "eu-de-2"
	instanceName := "test-instance"

	prototype := &vpcv1.InstancePrototypeInstanceBySourceTemplate{
		SourceTemplate: &vpcv1.InstanceTemplateIdentityByID{ID: &templateID},
		Zone:           &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Name:           &instanceName,
	}

	jsonBytes, err := json.Marshal(prototype)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	if !strings.Contains(jsonString, `"source_template"`) || !strings.Contains(jsonString, `"zone"`) {
		t.Errorf("InstancePrototypeInstanceBySourceTemplate missing interface fields in JSON: %s", jsonString)
	}
}

func TestInstancePrototypeInstanceByVolumeMarshal(t *testing.T) {
	volumeID := "r010-test-volume-id"
	zoneName := "eu-de-2"
	profileName := "bx2-2x8"
	vpcID := "r010-test-vpc-id"
	instanceName := "test-instance"

	prototype := &vpcv1.InstancePrototypeInstanceByVolume{
		BootVolumeAttachment: &vpcv1.VolumeAttachmentPrototypeInstanceByVolumeContext{
			Volume: &vpcv1.VolumeIdentityByID{ID: &volumeID},
		},
		Zone:    &vpcv1.ZoneIdentityByName{Name: &zoneName},
		Profile: &vpcv1.InstanceProfileIdentityByName{Name: &profileName},
		VPC:     &vpcv1.VPCIdentityByID{ID: &vpcID},
		Name:    &instanceName,
	}

	jsonBytes, err := json.Marshal(prototype)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonBytes)
	if !strings.Contains(jsonString, `"boot_volume_attachment"`) || !strings.Contains(jsonString, `"zone"`) ||
		!strings.Contains(jsonString, `"vpc"`) || !strings.Contains(jsonString, `"profile"`) {
		t.Errorf("InstancePrototypeInstanceByVolume missing interface fields in JSON: %s", jsonString)
	}
}
