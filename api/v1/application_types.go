/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	networkingv1alpha3 "istio.io/api/networking/v1alpha3"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ComponentsTemplate define the template of component
type ComponentsTemplate struct {
	Name    string        `json:"name"`
	Version string        `json:"version"`
	Spec    ComponentSpec `json:"spec,omitempty"`
}

// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	WorklordTemplate appsv1.DeploymentSpec `json:"template"`
	ComponentTraits  ComponentTraits       `json:"componentTraits,omitempty"`
}

//ComponentTraits defines the property of component
type ComponentTraits struct {
	CustomMetric *CustomMetric `json:"custommetric,omitempty"` //zk
	Logcollect   bool          `json:"logcollect,,omitempty"`
	Autoscaling  *Autoscaling  `json:"autoscaling,omitempty"` //zk
	Config       []ConfigFile  `json:"config,omitempty"`
}

//ConfigFile defines the property of config use for generate configmap and mount it
type ConfigFile struct {
	Path     string `json:"path"`
	FileName string `json:"fileName"`
	Value    string `json:"value"`
}

//Autoscaling defines the property of autoscaling use for generate hpa
type Autoscaling struct {
	Metric      string `json:"metric"`
	Threshold   int32  `json:"threshold"`
	MaxReplicas int32  `json:"maxreplicas"`
	MinReplicas int32  `json:"minreplicas"`
}

//CustomMetric defines the property of user costom metric
type CustomMetric struct {
	Enable bool   `json:"enable"`
	URI    string `json:"uri,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Application. Edit Application_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	Components []ComponentsTemplate `json:"components,omitempty"`
	Traits     AppTraits            `json:"appTraits,omitempty"`
}

// AppTraits defines the traits of Application
type AppTraits struct {
	ServiceSetting  *ServiceSetting  `json:"servicesetting,omitempty"`
	GrayRelease     map[string]int   `json:"grayRelease,omitempty"`
	ImagePullConfig *ImagePullConfig `json:"imagePullConfig,omitempty"`
}

//AppIngress defines the property of ingress
type AppIngress struct {
	Host       string `json:"host"`
	Path       string `json:"path,omitempty"`
	ServerPort int32  `json:"serverPort"`
}

//ServiceSetting defines the settings of Service
type ServiceSetting struct {
	TrafficPolicy *networkingv1alpha3.TrafficPolicy `json:"trafficpolicy,omitempty"`
	Ingress       AppIngress                        `json:"ingress,omitempty"`
	WhiteList     *WhiteList                        `json:"whiteList,omitempty"`
}

//ImagePullConfig defines the property of imagepullsecret
type ImagePullConfig struct {
	Registry string `json:"registry,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//WhiteList defines a list of users who can access the service
type WhiteList struct {
	Users []string `json:"users,omitempty"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
