/*
 * Copyright 2024 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import "time"

type IdentityType = string

type IdentityBase struct {
	Type     IdentityType   `json:"type"`
	Username string         `json:"username"`
	Meta     map[string]any `json:"meta"`
}

type Identity struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	IdentityBase
}

type IdentityFilter struct {
	Type IdentityType
}

type NewIdentityRequest struct {
	IdentityBase
	Secret string `json:"secret"`
}

type UpdateIdentityRequest struct {
	Meta   map[string]any `json:"meta"`
	Secret string         `json:"secret"`
}
