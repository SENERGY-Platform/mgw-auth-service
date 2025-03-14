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

package lib

import (
	"context"
	"github.com/SENERGY-Platform/mgw-auth-service/lib/model"
	srv_info_lib "github.com/SENERGY-Platform/mgw-go-service-base/srv-info-hdl/lib"
	"time"
)

type Api interface {
	GetIdentities(ctx context.Context, filter model.IdentityFilter) (map[string]model.Identity, error)
	GetIdentity(ctx context.Context, id string) (model.Identity, error)
	AddIdentity(ctx context.Context, base model.IdentityBase, secret string) (string, error)
	UpdateIdentity(ctx context.Context, id string, meta map[string]any, secret string) error
	DeleteIdentity(ctx context.Context, id string) error
	OpenPairingSession(ctx context.Context, duration time.Duration) error
	ClosePairingSession(ctx context.Context) error
	PairMachine(ctx context.Context, meta map[string]any) (cr model.CredentialsResponse, err error)
	srv_info_lib.Api
}
