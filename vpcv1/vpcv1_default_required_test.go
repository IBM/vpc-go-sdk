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
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

var _ = Describe("VpcV1Options struct invariants", func() {
	It("Version field must not be tagged with validate:\"required\"", func() {
		optionsType := reflect.TypeOf(vpcv1.VpcV1Options{})

		field, ok := optionsType.FieldByName("Version")
		Expect(ok).To(BeTrue(), "VpcV1Options must have a Version field")

		validateTag := field.Tag.Get("validate")
		// This will fail the test if someone adds validate:"required" back.
		Expect(validateTag).ToNot(ContainSubstring("required"),
			"VpcV1Options.Version must not be tagged with validate:\"required\"")
	})
})
