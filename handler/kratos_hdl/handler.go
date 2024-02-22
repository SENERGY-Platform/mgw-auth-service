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

package kratos_hdl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/go-service-base/context-hdl"
	lib_model "github.com/SENERGY-Platform/mgw-auth-service/lib/model"
	"github.com/SENERGY-Platform/mgw-auth-service/util"
	kratos "github.com/ory/kratos-client-go"
	"net/http"
	"time"
)

const (
	usernameKey = "username"
	metaKey     = "meta"
)

type errResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type Handler struct {
	kClient     *kratos.APIClient
	httpTimeout time.Duration
}

func New(apiClient *kratos.APIClient, httpTimeout time.Duration) *Handler {
	return &Handler{
		kClient:     apiClient,
		httpTimeout: httpTimeout,
	}
}

func (h *Handler) List(ctx context.Context, filter lib_model.IdentityFilter) (map[string]lib_model.Identity, error) {
	ctxWt, cf := context.WithTimeout(ctx, h.httpTimeout)
	defer cf()
	kIdentities, resp, err := h.kClient.IdentityAPI.ListIdentities(ctxWt).Execute()
	if err = handleResp(resp, err); err != nil {
		return nil, err
	}
	identities := make(map[string]lib_model.Identity)
	for _, kIdentity := range kIdentities {
		if filter.Type != "" && kIdentity.SchemaId != filter.Type {
			continue
		}
		identity, err := newIdentity(kIdentity)
		if err != nil {
			util.Logger.Error(err)
			continue
		}
		identities[kIdentity.Id] = identity
	}
	return identities, nil
}

func (h *Handler) Get(ctx context.Context, id string) (lib_model.Identity, error) {
	ctxWt, cf := context.WithTimeout(ctx, h.httpTimeout)
	defer cf()
	kIdentity, resp, err := h.kClient.IdentityAPI.GetIdentity(ctxWt, id).Execute()
	if err = handleResp(resp, err); err != nil {
		return lib_model.Identity{}, err
	}
	if kIdentity == nil {
		return lib_model.Identity{}, lib_model.NewInternalError(fmt.Errorf("request returned nil"))
	}
	identity, err := newIdentity(*kIdentity)
	if err != nil {
		return lib_model.Identity{}, lib_model.NewInternalError(err)
	}
	return identity, nil
}

func (h *Handler) Add(ctx context.Context, iBase lib_model.IdentityBase, secret string) (string, error) {
	ctxWt, cf := context.WithTimeout(ctx, h.httpTimeout)
	defer cf()
	config := kratos.IdentityWithCredentialsPasswordConfig{}
	config.SetPassword(secret)
	state := "active"
	body := kratos.CreateIdentityBody{
		Credentials: &kratos.IdentityWithCredentials{
			Password: &kratos.IdentityWithCredentialsPassword{
				Config: &config,
			},
		},
		SchemaId: iBase.Type,
		State:    &state,
		Traits:   newTraits(iBase.Username, iBase.Meta),
	}
	kIdent, resp, err := h.kClient.IdentityAPI.CreateIdentity(ctxWt).CreateIdentityBody(body).Execute()
	if err = handleResp(resp, err); err != nil {
		return "", err
	}
	if kIdent == nil {
		return "", lib_model.NewInternalError(fmt.Errorf("request returned nil"))
	}
	return kIdent.Id, nil
}

func (h *Handler) Update(ctx context.Context, id string, meta map[string]any, secret string) error {
	ch := context_hdl.New()
	defer ch.CancelAll()
	kIdentity, resp, err := h.kClient.IdentityAPI.GetIdentity(ch.Add(context.WithTimeout(ctx, h.httpTimeout)), id).Execute()
	if err = handleResp(resp, err); err != nil {
		return err
	}
	username, _, err := getTraits(kIdentity.Traits)
	if err != nil {
		return fmt.Errorf("error getting identiy '%s' traits: '%s'", kIdentity.Id, err)
	}
	config := kratos.IdentityWithCredentialsPasswordConfig{}
	config.SetPassword(secret)
	body := kratos.UpdateIdentityBody{
		Credentials: &kratos.IdentityWithCredentials{
			Password: &kratos.IdentityWithCredentialsPassword{
				Config: &config,
			},
		},
		SchemaId: kIdentity.SchemaId,
		Traits:   newTraits(username, meta),
	}
	if kIdentity.State != nil {
		body.State = *kIdentity.State
	}
	_, resp, err = h.kClient.IdentityAPI.UpdateIdentity(ch.Add(context.WithTimeout(ctx, h.httpTimeout)), id).UpdateIdentityBody(body).Execute()
	if err = handleResp(resp, err); err != nil {
		return err
	}
	resp, err = h.kClient.IdentityAPI.DeleteIdentitySessions(ch.Add(context.WithTimeout(ctx, h.httpTimeout)), id).Execute()
	if err = handleResp(resp, err); err != nil {
		return err
	}
	return nil
}

func (h *Handler) Delete(ctx context.Context, id string) error {
	ctxWt, cf := context.WithTimeout(ctx, h.httpTimeout)
	defer cf()
	resp, err := h.kClient.IdentityAPI.DeleteIdentity(ctxWt, id).Execute()
	if err = handleResp(resp, err); err != nil {
		return err
	}
	return nil
}

func getTraits(traits any) (string, map[string]any, error) {
	traitMap, ok := traits.(map[string]any)
	if !ok {
		return "", nil, errors.New("traits invalid type")
	}
	var username string
	u, ok := traitMap[usernameKey]
	if !ok {
		return "", nil, errors.New("missing username")
	}
	if username, ok = u.(string); !ok {
		return "", nil, errors.New("username invalid type")
	}
	var meta map[string]any
	m, ok := traitMap[metaKey]
	if ok {
		if meta, ok = m.(map[string]any); !ok {
			return "", nil, errors.New("meta invalid type")
		}
	}
	return username, meta, nil
}

func newIdentity(kIdentity kratos.Identity) (lib_model.Identity, error) {
	username, meta, err := getTraits(kIdentity.Traits)
	if err != nil {
		return lib_model.Identity{}, fmt.Errorf("error getting identiy '%s' traits: '%s'", kIdentity.Id, err)
	}
	identity := lib_model.Identity{
		ID: kIdentity.Id,
		IdentityBase: lib_model.IdentityBase{
			Type:     kIdentity.SchemaId,
			Username: username,
			Meta:     meta,
		},
	}
	if kIdentity.CreatedAt != nil {
		identity.Created = *kIdentity.CreatedAt
	}
	if kIdentity.UpdatedAt != nil {
		identity.Updated = *kIdentity.UpdatedAt
	}
	return identity, nil
}

func handleResp(resp *http.Response, err error) error {
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		if resp != nil {
			decoder := json.NewDecoder(resp.Body)
			var errResp errResponse
			e := decoder.Decode(&errResp)
			if e == nil {
				err = errors.New(errResp.Error.Message)
				if resp.StatusCode == http.StatusNotFound {
					return lib_model.NewNotFoundError(err)
				}
			}
		}
		return lib_model.NewInternalError(err)
	}
	return nil
}

func newTraits(username string, meta map[string]any) map[string]any {
	return map[string]any{
		usernameKey: username,
		metaKey:     meta,
	}
}
