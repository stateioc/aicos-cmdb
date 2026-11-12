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

package topo_server_test

import (
	"testing"

	"configcenter/src/test"
	"configcenter/src/test/reporter"
	"configcenter/src/test/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var header = test.GetHeader()
var clientSet = test.GetClientSet()
var topoServerClient = clientSet.TopoServer()
var procServerClient = clientSet.ProcServer()
var apiServerClient = clientSet.ApiServer()
var instClient = topoServerClient.Instance()
var asstClient = topoServerClient.Association()
var objectClient = topoServerClient.Object()
var serviceClient = clientSet.ProcServer().Service()
var hostServerClient = test.GetClientSet().HostServer()

var kubeClient = topoServerClient.Kube()
var fieldTemplateClient = topoServerClient.FieldTemplate()
var modelQuoteClient = apiServerClient.ModelQuote()

func TestTopoServer(t *testing.T) {
	RegisterFailHandler(util.Fail)
	reporters := []Reporter{
		reporter.NewHtmlReporter(test.GetReportDir()+"toposerver.html", test.GetReportUrl(), true),
	}
	RunSpecsWithDefaultAndCustomReporters(t, "TopoServer Suite", reporters)
}

var _ = BeforeSuite(func() {
	test.ClearDatabase()
})
