// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bastion

import (
	"testing"

	"github.com/gardener/gardener/extensions/pkg/controller"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/extensions"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestBastion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bastion Suite")
}

var _ = Describe("Bastion", func() {
	var (
		cluster *extensions.Cluster
		bastion *extensionsv1alpha1.Bastion

		ctrl                 *gomock.Controller
		maxLengthForResource int
	)
	BeforeEach(func() {
		cluster = createOpenstackTestCluster()
		bastion = createTestBastion()
		ctrl = gomock.NewController(GinkgoT())
		maxLengthForResource = 63
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Determine options", func() {
		It("should return options", func() {
			options, err := DetermineOptions(bastion, cluster)
			Expect(err).To(Not(HaveOccurred()))

			Expect(options.ShootName).To(Equal("cluster1"))
			Expect(options.BastionInstanceName).To(Equal("cluster1-bastionName1-bastion-1cdc8"))
			Expect(options.SecretReference).To(Equal(corev1.SecretReference{
				Namespace: "cluster1",
				Name:      "cloudprovider",
			}))
			Expect(options.Region).To(Equal("eu-nl-1"))
			Expect(options.FloatingPoolName).To(Equal("FloatingIP-external-monsoon-testing"))

		})
	})

	Describe("check Ingress Permissions", func() {
		It("Should return a string array with ipV4 normalized addresses", func() {
			bastion.Spec.Ingress = []extensionsv1alpha1.BastionIngressPolicy{
				{IPBlock: networkingv1.IPBlock{
					CIDR: "0.0.0.0/0",
				}},
			}
			etherType, _, err := ingressPermissions(bastion)
			Expect(err).To(Not(HaveOccurred()))
			Expect(string(etherType)).To(Equal("IPv4"))
		})
		It("Should return a string array with ipV6 normalized addresses", func() {
			bastion.Spec.Ingress = []extensionsv1alpha1.BastionIngressPolicy{
				{IPBlock: networkingv1.IPBlock{
					CIDR: "::/0",
				}},
			}
			etherType, _, err := ingressPermissions(bastion)
			Expect(err).To(Not(HaveOccurred()))
			Expect(string(etherType)).To(Equal("IPv6"))
		})
	})

	Describe("check Names generations", func() {
		It("should generate idempotent name", func() {
			expected := "clusterName-shortName-bastion-79641"

			res, err := generateBastionBaseResourceName("clusterName", "shortName")
			Expect(err).To(Not(HaveOccurred()))
			Expect(res).To(Equal(expected))

			res, err = generateBastionBaseResourceName("clusterName", "shortName")
			Expect(err).To(Not(HaveOccurred()))
			Expect(res).To(Equal(expected))
		})

		It("should generate a name not exceeding a certain length", func() {
			res, err := generateBastionBaseResourceName("clusterName", "LetsExceed63LenLimit012345678901234567890123456789012345678901234567890123456789")
			Expect(err).To(Not(HaveOccurred()))
			Expect(res).To(Equal("clusterName-LetsExceed63LenLimit0-bastion-139c4"))
		})

		It("should generate a unique name even if inputs values have minor deviations", func() {
			res, _ := generateBastionBaseResourceName("1", "1")
			res2, _ := generateBastionBaseResourceName("1", "2")
			Expect(res).ToNot(Equal(res2))
		})

		baseName, _ := generateBastionBaseResourceName("clusterName", "LetsExceed63LenLimit012345678901234567890123456789012345678901234567890123456789")
		DescribeTable("should generate names and fit maximum length",
			func(input string, expectedOut string) {
				Expect(len(input)).Should(BeNumerically("<", maxLengthForResource))
				Expect(input).Should(Equal(expectedOut))
			},

			Entry("security group name", securityGroupName(baseName), "clusterName-LetsExceed63LenLimit0-bastion-139c4-sg"),
			Entry("ingress allow ssh resource name", ingressAllowSSHResourceName(baseName), "clusterName-LetsExceed63LenLimit0-bastion-139c4-allow-ssh"),
			Entry("egress allow only resource name", egressAllowOnlyResourceName(baseName), "clusterName-LetsExceed63LenLimit0-bastion-139c4-egress-worker"),
		)
	})

})

func createTestBastion() *extensionsv1alpha1.Bastion {
	return &extensionsv1alpha1.Bastion{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bastionName1",
		},
		Spec: extensionsv1alpha1.BastionSpec{
			DefaultSpec: extensionsv1alpha1.DefaultSpec{},
			UserData:    nil,
			Ingress: []extensionsv1alpha1.BastionIngressPolicy{
				{IPBlock: networkingv1.IPBlock{
					CIDR: "213.69.151.0/24",
				}},
			},
		},
	}
}

func createOpenstackTestCluster() *extensions.Cluster {
	return &controller.Cluster{
		ObjectMeta: metav1.ObjectMeta{Name: "cluster1"},
		Shoot:      createShootTestStruct(),
		CloudProfile: &gardencorev1beta1.CloudProfile{
			Spec: gardencorev1beta1.CloudProfileSpec{
				Regions: []gardencorev1beta1.Region{
					{Name: ("eu-nl-1")},
				},
			},
		},
	}
}

func createShootTestStruct() *gardencorev1beta1.Shoot {
	json := `{"apiVersion": "openstack.provider.extensions.gardener.cloud/v1alpha1","kind": "InfrastructureConfig", "FloatingPoolName": "FloatingIP-external-monsoon-testing"}`
	return &gardencorev1beta1.Shoot{
		Spec: gardencorev1beta1.ShootSpec{
			Region:            "eu-nl-1",
			SecretBindingName: v1beta1constants.SecretNameCloudProvider,
			Provider: gardencorev1beta1.Provider{
				InfrastructureConfig: &runtime.RawExtension{
					Raw: []byte(json),
				}}},
	}
}
