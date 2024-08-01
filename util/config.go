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
	"github.com/SENERGY-Platform/go-service-base/config-hdl"
	cfg_types "github.com/SENERGY-Platform/go-service-base/config-hdl/types"
	sb_logger "github.com/SENERGY-Platform/go-service-base/logger"
	envldr "github.com/y-du/go-env-loader"
	"github.com/y-du/go-log-level/level"
	"reflect"
)

type HttpClientConfig struct {
	IdentitySrvBaseUrl string `json:"identity_srv_base_url" env_var:"IDENTITY_SRV_BASE_URL"`
	Timeout            int64  `json:"timeout" env_var:"HTTP_TIMEOUT"`
}

type InitIdentityConfig struct {
	User   string           `json:"user" env_var:"II_USER"`
	Secret cfg_types.Secret `json:"secret" env_var:"II_SECRET"`
}

type LoggerConfig struct {
	Level        level.Level `json:"level" env_var:"LOGGER_LEVEL"`
	Utc          bool        `json:"utc" env_var:"LOGGER_UTC"`
	Path         string      `json:"path" env_var:"LOGGER_PATH"`
	FileName     string      `json:"file_name" env_var:"LOGGER_FILE_NAME"`
	Terminal     bool        `json:"terminal" env_var:"LOGGER_TERMINAL"`
	Microseconds bool        `json:"microseconds" env_var:"LOGGER_MICROSECONDS"`
	Prefix       string      `json:"prefix" env_var:"LOGGER_PREFIX"`
}

type Config struct {
	ServerPort    uint               `json:"server_port" env_var:"SERVER_PORT"`
	Logger        LoggerConfig       `json:"logger" env_var:"LOGGER_CONFIG"`
	HttpClient    HttpClientConfig   `json:"http_client" env_var:"HTTP_CLIENT_CONFIG"`
	CSDefDuration int64              `json:"cs_def_duration" env_var:"CS_DEF_DURATION"`
	InitIdentity  InitIdentityConfig `json:"init_identity" env_var:"INIT_IDENTITY"`
}

func NewConfig(path string) (*Config, error) {
	cfg := Config{
		ServerPort: 80,
		Logger: LoggerConfig{
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
	err := config_hdl.Load(&cfg, nil, map[reflect.Type]envldr.Parser{reflect.TypeOf(level.Off): sb_logger.LevelParser}, nil, path)
	return &cfg, err
}
