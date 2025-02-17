package jenkinsgithub

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"

	"github.com/devstream-io/devstream/pkg/util/k8s"
	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/devstream-io/devstream/pkg/util/template"
)

const githubIntegName = "github-integ"

//go:embed tpl/github-integ.tpl.yaml
var githubIntegTemplate string

func applyGitHubIntegConfig(opts *Options) error {
	githubIntegOption := &GitHubIntegOptions{
		AdminList:          opts.AdminList,
		CredentialsID:      jenkinsCredentialID,
		JenkinsURLOverride: opts.J.URLOverride,
		GithubAuthID:       githubAuthID,
	}

	gitHubIntegContent, err := renderGitHubInteg(githubIntegOption)
	if err != nil {
		return fmt.Errorf("failed to render github integ JCasC: %s", err)
	}

	if err := applyJCasC(opts.Helm.Namespace, opts.Helm.ReleaseName, githubIntegName, gitHubIntegContent); err != nil {
		return fmt.Errorf("failed to create config map for github integ JCasC: %s", err)
	}

	return nil
}

// refer: https://github.com/jenkinsci/helm-charts/blob/main/charts/jenkins/templates/jcasc-config.yaml#L6-L25
func applyJCasC(namespace, chartReleaseName, configName, fileContent string) error {
	client, err := k8s.NewClient()
	if err != nil {
		return err
	}

	configMapName := fmt.Sprintf("%s-jenkins-jenkins-config-%s", chartReleaseName, configName)
	labels := map[string]string{
		"app.kubernetes.io/component":                              "jenkins-controller",
		"app.kubernetes.io/instance":                               chartReleaseName,
		"app.kubernetes.io/managed-by":                             "DevStream",
		"created_by":                                               "DevStream",
		fmt.Sprintf("%s-jenkins-jenkins-config", chartReleaseName): "true",
	}
	data := map[string]string{
		// relaseName-chartName-jenkins-config
		fmt.Sprintf("%s.yaml", configName): fileContent,
	}

	configMap := corev1.ConfigMap(configMapName, namespace).
		WithLabels(labels).
		WithData(data).
		WithImmutable(false)

	applyOptions := metav1.ApplyOptions{
		FieldManager: "DevStream",
	}

	configMapRes, err := client.CoreV1().ConfigMaps(namespace).Apply(context.TODO(), configMap, applyOptions)
	if err != nil {
		return err
	}

	log.Debugf("Created configmap %+v", configMapRes)

	// wait for the config map and the sidecar to be ready
	// TODO(aFlyBird0): read JCasC to judge if JCasC is ready
	time.Sleep(time.Second * 3)

	return nil
}

func renderGitHubInteg(opts *GitHubIntegOptions) (string, error) {
	return template.New().FromContent(githubIntegTemplate).SetDefaultRender(githubIntegName, opts).Render()
}
