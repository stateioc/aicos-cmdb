/*
 * SPACE-AICOS is pleased to support the open source community by making
 * SPACE-AICOS available.
 * Copyright (C) 2026 China Academy of Information and Communications Technology (CAICT).
 * All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package healthz_test

import (
	"testing"

	"configcenter/src/common/types"
	"configcenter/src/test"
	"configcenter/src/test/reporter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var header = test.GetHeader()
var clientSet = test.GetClientSet()
var healthzClient = clientSet.Healthz()

func TestHealthz(t *testing.T) {
	RegisterFailHandler(Fail)
	reporters := []Reporter{
		reporter.NewHtmlReporter(test.GetReportDir()+"healthz.html", test.GetReportUrl(), true),
	}
	RunSpecsWithDefaultAndCustomReporters(t, "Healthz Suite", reporters)
}

var _ = Describe("healthz test", func() {
	It(types.CC_MODULE_DATACOLLECTION+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_DATACOLLECTION)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_HOST+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_HOST)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_MIGRATE+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_MIGRATE)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_PROC+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_PROC)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_TOPO+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_TOPO)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_EVENTSERVER+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_EVENTSERVER)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_APISERVER+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_APISERVER)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})

	It(types.CC_MODULE_CORESERVICE+" healthz test", func() {
		isHealthy, err := healthzClient.HealthCheck(types.CC_MODULE_CORESERVICE)
		Expect(err).NotTo(HaveOccurred())
		Expect(isHealthy).To(Equal(true))
	})
})
