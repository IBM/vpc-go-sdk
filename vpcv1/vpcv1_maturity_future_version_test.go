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
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

var _ = Describe("VpcV1 Maturity and FutureVersion Tests", func() {
	var (
		testServer *httptest.Server
		vpcService *vpcv1.VpcV1
	)

	BeforeEach(func() {
		testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"vpcs": []}`)
		}))

		var err error
		vpcService, err = vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
			URL:           testServer.URL,
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(err).To(BeNil())
		Expect(vpcService).ToNot(BeNil())
	})

	AfterEach(func() {
		testServer.Close()
	})

	Context("Maturity field", func() {
		It("should set maturity query parameter when Maturity is set to 'development'", func() {
			maturity := "development"
			vpcService.Maturity = &maturity

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("maturity")).To(Equal("development"))
		})

		It("should not set maturity query parameter when Maturity is nil", func() {
			vpcService.Maturity = nil

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("maturity")).To(Equal(""))
		})

		It("should not set maturity query parameter when Maturity is empty string", func() {
			maturity := ""
			vpcService.Maturity = &maturity

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("maturity")).To(Equal(""))
		})
	})

	Context("FutureVersion field", func() {
		It("should set future_version query parameter when FutureVersion is set", func() {
			futureVersion := "true"
			vpcService.FutureVersion = &futureVersion

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("future_version")).To(Equal("true"))
		})

		It("should not set future_version query parameter when FutureVersion is nil", func() {
			vpcService.FutureVersion = nil

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("future_version")).To(Equal(""))
		})

		It("should not set future_version query parameter when FutureVersion is empty string", func() {
			futureVersion := ""
			vpcService.FutureVersion = &futureVersion

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("future_version")).To(Equal(""))
		})

		It("should set future_version but NOT override version when FutureVersion is 'true'", func() {
			futureVersion := "true"
			vpcService.FutureVersion = &futureVersion

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()

			// Verify future_version is set
			Expect(queryParams.Get("future_version")).To(Equal("true"))

			// When FutureVersion is "true", version should NOT be overridden to tomorrow
			// It should keep the original version parameter
			Expect(queryParams.Has("version")).To(BeTrue())
		})

		It("should override version to tomorrow's date when FutureVersion is not 'true'", func() {
			futureVersion := "false"
			vpcService.FutureVersion = &futureVersion

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()

			// Verify future_version is set
			Expect(queryParams.Get("future_version")).To(Equal("false"))

			// Verify version is set to tomorrow's date
			tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
			Expect(queryParams.Get("version")).To(Equal(tomorrow))
		})
	})

	Context("Combined Maturity and FutureVersion", func() {
		It("should set both maturity and future_version query parameters when both are set", func() {
			maturity := "development"
			futureVersion := "true"
			vpcService.Maturity = &maturity
			vpcService.FutureVersion = &futureVersion

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()
			Expect(queryParams.Get("maturity")).To(Equal("development"))
			Expect(queryParams.Get("future_version")).To(Equal("true"))
		})

		It("should set maturity and override version when FutureVersion is not 'true'", func() {
			maturity := "development"
			futureVersion := "2026-01-01"
			vpcService.Maturity = &maturity
			vpcService.FutureVersion = &futureVersion

			var capturedRequest *http.Request
			testServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"vpcs": []}`)
			})

			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			_, _, err := vpcService.ListVpcs(listVpcsOptions)
			Expect(err).To(BeNil())

			Expect(capturedRequest).ToNot(BeNil())
			queryParams := capturedRequest.URL.Query()

			// Verify both parameters are set
			Expect(queryParams.Get("maturity")).To(Equal("development"))
			Expect(queryParams.Get("future_version")).To(Equal("2026-01-01"))

			// Verify version is set to tomorrow's date
			tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
			Expect(queryParams.Get("version")).To(Equal(tomorrow))
		})
	})
})

// Made with Bob
