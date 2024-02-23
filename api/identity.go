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

package api

import (
	"context"
	lib_model "github.com/SENERGY-Platform/mgw-auth-service/lib/model"
	"time"
)

func (a *Api) GetIdentities(ctx context.Context, filter lib_model.IdentityFilter) (map[string]lib_model.Identity, error) {
	return a.identityHdl.List(ctx, filter)
}

func (a *Api) GetIdentity(ctx context.Context, id string) (lib_model.Identity, error) {
	return a.identityHdl.Get(ctx, id)
}

func (a *Api) AddIdentity(ctx context.Context, base lib_model.IdentityBase, secret string) (string, error) {
	return a.identityHdl.Add(ctx, base, secret)
}

func (a *Api) UpdateIdentity(ctx context.Context, id string, meta map[string]any, secret string) error {
	return a.identityHdl.Update(ctx, id, meta, secret)
}

func (a *Api) DeleteIdentity(ctx context.Context, id string) error {
	return a.identityHdl.Delete(ctx, id)
}

func (a *Api) OpenPairingSession(_ context.Context, duration time.Duration) error {
	return a.credentialSessionHdl.Open(duration)
}

func (a *Api) ClosePairingSession(_ context.Context) error {
	a.credentialSessionHdl.Close()
	return nil
}

func (a *Api) PairMachine(ctx context.Context, meta map[string]any) (lib_model.CredentialsResponse, error) {
	var cr lib_model.CredentialsResponse
	var err error
	cr.Login, cr.Secret, err = a.credentialSessionHdl.GetCredentials()
	if err != nil {
		return lib_model.CredentialsResponse{}, err
	}
	base := lib_model.IdentityBase{
		Type:     lib_model.MachineType,
		Username: cr.Login,
		Meta:     meta,
	}
	cr.ID, err = a.identityHdl.Add(ctx, base, cr.Secret)
	if err != nil {
		return lib_model.CredentialsResponse{}, err
	}
	return cr, nil
}
