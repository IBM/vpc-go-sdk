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
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

var _ = Describe("NewVpcV1 Version invariants", func() {
	It("must always return a service with a non-nil Version in YYYY-MM-DD format", func() {
		auth := &core.NoAuthAuthenticator{}

		opts := &vpcv1.VpcV1Options{
			Authenticator: auth,
			// Version intentionally omitted to exercise defaulting.
		}

		svc, err := vpcv1.NewVpcV1(opts)
		Expect(err).To(BeNil(), "NewVpcV1 should not fail when Version is omitted")
		Expect(svc).ToNot(BeNil(), "NewVpcV1 must return a non-nil service")

		// Invariant: service.Version must never be nil or empty.
		Expect(svc.Version).ToNot(BeNil(), "service.Version must not be nil")

		version := *svc.Version
		Expect(version).ToNot(BeEmpty(), "service.Version must not be empty")

		// Invariant: Version must be in strict YYYY-MM-DD format.
		re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
		Expect(re.MatchString(version)).To(BeTrue(),
			"service.Version must be in strict YYYY-MM-DD format, got: %s", version)
	})
})
