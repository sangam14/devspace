package v1beta2

import (
	"reflect"
	"testing"

	next "github.com/devspace-cloud/devspace/pkg/devspace/config/versions/latest"
	"github.com/devspace-cloud/devspace/pkg/util/ptr"
	yaml "gopkg.in/yaml.v2"
)

type testCase struct {
	in       *Config
	expected *next.Config
}

func TestSimple(t *testing.T) {
	testCases := []*testCase{
		{
			in: &Config{
				Cluster: &Cluster{
					Namespace:   ptr.String("namespace"),
					KubeContext: ptr.String("kubecontext"),
				},
			},
			expected: &next.Config{
				Dev: &next.DevConfig{
					Terminal: &next.Terminal{
						Enabled: ptr.Bool(true),
					},
				},
			},
		},
		{
			in: &Config{
				Dev: &DevConfig{
					Terminal: &Terminal{
						Disabled: ptr.Bool(true),
					},
				},
				Images: &map[string]*ImageConfig{
					"default": &ImageConfig{},
				},
			},
			expected: &next.Config{
				Dev: &next.DevConfig{
					Terminal: &next.Terminal{
						Enabled: ptr.Bool(false),
					},
				},
				Images: map[string]*next.ImageConfig{
					"default": &next.ImageConfig{
						CreatePullSecret: ptr.Bool(false),
					},
				},
			},
		},
	}

	// Run test cases
	for index, testCase := range testCases {
		newConfig, err := testCase.in.Upgrade()
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		isEqual := reflect.DeepEqual(newConfig, testCase.expected)
		if !isEqual {
			newConfigYaml, _ := yaml.Marshal(newConfig)
			expectedYaml, _ := yaml.Marshal(testCase.expected)

			t.Fatalf("TestCase %d: Got %s, but expected %s", index, newConfigYaml, expectedYaml)
		}
	}
}
