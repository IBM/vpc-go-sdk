/**
 * (C) Copyright IBM Corp. 2026.
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

package vpcv1

import (
	"regexp"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("getServiceComponentInfo function", func() {
	It("Should return ProblemComponent with DefaultServiceName and date in YYYY-MM-DD format", func() {
		// Call the function directly since we're in the same package
		component := getServiceComponentInfo()

		// Verify the component is not nil
		Expect(component).ToNot(BeNil())

		// Verify the component name matches DefaultServiceName
		Expect(component.Name).To(Equal(DefaultServiceName),
			"Component Name should be DefaultServiceName (vpc)")

		// Verify the version is in YYYY-MM-DD format
		datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
		Expect(datePattern.MatchString(component.Version)).To(BeTrue(),
			"Version should be in YYYY-MM-DD format, got: %s", component.Version)
	})

	It("Should return the expected ProblemComponent structure", func() {
		component := getServiceComponentInfo()

		// Verify it's a valid ProblemComponent pointer
		Expect(component).To(BeAssignableToTypeOf(&core.ProblemComponent{}))
	})
})
