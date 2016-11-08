/*
Copyright 2016 The Kubernetes Authors.

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

package dockerhub

// CallbackBody encapsulates DockerHub webhook callback request bodies
type CallbackBody struct {
	CallbackURL string     `json:"callback_url,omitempty"`
	PushData    PushData   `json:"push_data,omitempty"`
	Repository  Repository `json:",omitempty"`
}

// PushData encapsulates data about the pushed image
type PushData struct {
	Images   []string `json:images",omitempty"`
	PushedAt int      `json:"pushed_at,omitempty"` // Required
	Pusher   string   `json:pusher",omitempty"`
	Tag      string   `json:tag",omitempty"` // Required
}

// Repository encapsulates data about the dockerhub repository
type Repository struct {
	DateCreated     int    `json:"date_created,omitempty"`
	Description     string `json:"description,omitempty"`
	IsTrusted       bool   `json:"is_trusted,omitempty"`
	FullDescription string `json:"full_description,omitempty"`
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Owner           string `json:"owner,omitempty"`
	RepoName        string `json:"repo_name,omitempty"` // Required
	RepoURL         string `json:"repo_url,omitempty"`
	Status          string `json:"status,omitempty"`
}
