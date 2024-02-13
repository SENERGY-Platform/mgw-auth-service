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

type DatabaseConfig struct {
	Host       string               `json:"host" env_var:"DB_HOST"`
	Port       uint                 `json:"port" env_var:"DB_PORT"`
	User       string               `json:"user" env_var:"DB_USER"`
	Passwd     sb_util.SecretString `json:"passwd" env_var:"DB_PASSWD"`
	Name       string               `json:"name" env_var:"DB_NAME"`
	Timeout    int64                `json:"timeout" env_var:"DB_TIMEOUT"`
	SchemaPath string               `json:"schema_path" env_var:"DB_SCHEMA_PATH"`
}

type Config struct {
	ServerPort uint                 `json:"server_port" env_var:"SERVER_PORT"`
	Logger     sb_util.LoggerConfig `json:"logger" env_var:"LOGGER_CONFIG"`
	Database   DatabaseConfig       `json:"database" env_var:"DATABASE_CONFIG"`
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
		Database: DatabaseConfig{
			Host:       "core-db",
			Port:       3306,
			Name:       "auth_service",
			Timeout:    5000000000,
			SchemaPath: "include/auth_storage_schema.sql",
		},
	}
	err := sb_util.LoadConfig(path, &cfg, nil, nil, nil)
	return &cfg, err
}
