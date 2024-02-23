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

package util

import (
	sb_util "github.com/SENERGY-Platform/go-service-base/util"
	"github.com/y-du/go-log-level/level"
)

type HttpClientConfig struct {
	IdentitySrvBaseUrl string `json:"identity_srv_base_url" env_var:"IDENTITY_SRV_BASE_URL"`
	Timeout            int64  `json:"timeout" env_var:"HTTP_TIMEOUT"`
}

type InitIdentityConfig struct {
	User   string               `json:"user" env_var:"II_USER"`
	Secret sb_util.SecretString `json:"secret" env_var:"II_SECRET"`
}

type Config struct {
	ServerPort    uint                 `json:"server_port" env_var:"SERVER_PORT"`
	Logger        sb_util.LoggerConfig `json:"logger" env_var:"LOGGER_CONFIG"`
	HttpClient    HttpClientConfig     `json:"http_client" env_var:"HTTP_CLIENT_CONFIG"`
	CSDefDuration int64                `json:"cs_def_duration" env_var:"CS_DEF_DURATION"`
	InitIdentity  InitIdentityConfig   `json:"init_identity" env_var:"INIT_IDENTITY"`
}

func NewConfig(path string) (*Config, error) {
	cfg := Config{
		ServerPort: 80,
		Logger: sb_util.LoggerConfig{
			Level:        level.Warning,
			Utc:          true,
			Microseconds: true,
			Terminal:     true,
		},
		HttpClient: HttpClientConfig{
			IdentitySrvBaseUrl: "http://identity-service",
			Timeout:            10000000000,
		},
		CSDefDuration: 300000000000,
	}
	err := sb_util.LoadConfig(path, &cfg, nil, nil, nil)
	return &cfg, err
}
