package controllers

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	. "github.com/onsi/ginkgo"
	"net/http"
	"net/http/httputil"
	"time"
)

var _ = Describe("Contracttests controller", func() {

	const (
		namespace = "default"
		name      = "test-contract"
	)

	//log := logf.Log.WithName("ContractTests controller test")

	Context("When contracttest is created without timeout", func() {
		It("trying http requests", func() {
			time.Sleep(time.Second * 20)
			fmt.Println(testEnv.Config)
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(testEnv.Config.CAData)
			request, err := http.NewRequest("GET", testEnv.Config.Host+"api/v1/namespaces", nil)
			if err != nil {
				return
			}
			request.Header = http.Header{
				"Authorization": []string{"Bearer " + testEnv.Config.BearerToken},
			}
			client := &http.Client{

				Transport: &http.Transport{

					TLSClientConfig: &tls.Config{
						RootCAs: caCertPool,
					},
				},
			}
			get, err := client.Do(request)
			if err != nil {
				fmt.Println(err)
				return
			}
			response, err := httputil.DumpResponse(get, true)
			if err != nil {
				return
			}
			fmt.Println(string(response))
		})
		//It("should greet immediately in status", func() {
		//
		//	By("By creating new ContractTest CR")
		//	ctx := context.Background()
		//	contracttest := &webappv1.ContractTests{
		//		TypeMeta: metav1.TypeMeta{
		//			Kind:       "ContractTest",
		//			APIVersion: "webapp.appstudio.qe/v1",
		//		},
		//		ObjectMeta: metav1.ObjectMeta{
		//			Name:      name,
		//			Namespace: namespace,
		//		},
		//		Spec: webappv1.ContractTestsSpec{
		//			ContractName: "Test",
		//			WaitSecs:     3,
		//		},
		//	}
		//	Expect(k8sClient.Create(ctx, contracttest)).Should(Succeed())
		//
		//	contractTestLookupKey := types.NamespacedName{
		//		Namespace: namespace,
		//		Name:      name,
		//	}
		//	createdContractTest := &webappv1.ContractTests{}
		//
		//	Eventually(func() (string, error) {
		//		log.Info("Polling Get")
		//		err := k8sClient.Get(ctx, contractTestLookupKey, createdContractTest)
		//		if err != nil {
		//			return "", err
		//		}
		//		return createdContractTest.Status.Message, nil
		//	}, time.Second*10, time.Millisecond*250).Should(Equal("Hello Test"))
		//})
	})
})
