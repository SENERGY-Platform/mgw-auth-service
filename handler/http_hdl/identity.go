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

package http_hdl

import (
	"github.com/SENERGY-Platform/mgw-auth-service/lib"
	lib_model "github.com/SENERGY-Platform/mgw-auth-service/lib/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type getIdentitiesQuery struct {
	Type string `form:"type"`
}

func GetIdentitiesH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, lib_model.IdentitiesPath, func(gc *gin.Context) {
		query := getIdentitiesQuery{}
		if err := gc.ShouldBindQuery(&query); err != nil {
			_ = gc.Error(lib_model.NewInvalidInputError(err))
			return
		}
		filter := lib_model.IdentityFilter{
			Type: query.Type,
		}
		identities, err := a.GetIdentities(gc.Request.Context(), filter)
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.JSON(http.StatusOK, identities)
	}
}

func PostIdentityH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, lib_model.IdentitiesPath, func(gc *gin.Context) {
		var req lib_model.NewIdentityRequest
		err := gc.ShouldBindJSON(&req)
		if err != nil {
			_ = gc.Error(lib_model.NewInvalidInputError(err))
			return
		}
		id, err := a.AddIdentity(gc.Request.Context(), req.IdentityBase, req.Secret)
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.String(http.StatusOK, id)
	}
}

func GetIdentityH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, path.Join(lib_model.IdentitiesPath, ":id"), func(gc *gin.Context) {
		identity, err := a.GetIdentity(gc.Request.Context(), gc.Param("id"))
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.JSON(http.StatusOK, identity)
	}
}

func PatchIdentityH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPatch, path.Join(lib_model.IdentitiesPath, ":id"), func(gc *gin.Context) {
		var req lib_model.UpdateIdentityRequest
		err := gc.ShouldBindJSON(&req)
		if err != nil {
			_ = gc.Error(lib_model.NewInvalidInputError(err))
			return
		}
		err = a.UpdateIdentity(gc.Request.Context(), gc.Param("id"), req.Meta, req.Secret)
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.Status(http.StatusOK)
	}
}

func DeleteIdentityH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodDelete, path.Join(lib_model.IdentitiesPath, ":id"), func(gc *gin.Context) {
		if err := a.DeleteIdentity(gc.Request.Context(), gc.Param("id")); err != nil {
			_ = gc.Error(err)
			return
		}
		gc.Status(http.StatusOK)
	}
}
