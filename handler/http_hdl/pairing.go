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
	"time"
)

type pairingOpenQuery struct {
	Duration int64 `form:"duration"`
}

func postPairingH(a lib.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		var req map[string]any
		err := gc.ShouldBindJSON(&req)
		if err != nil {
			_ = gc.Error(lib_model.NewInvalidInputError(err))
			return
		}
		cr, err := a.PairMachine(gc.Request.Context(), req)
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.JSON(http.StatusOK, cr)
	}
}

func patchPairingOpenH(a lib.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		var query pairingOpenQuery
		if err := gc.ShouldBindQuery(&query); err != nil {
			_ = gc.Error(lib_model.NewInvalidInputError(err))
			return
		}
		err := a.OpenPairingSession(gc.Request.Context(), time.Duration(query.Duration))
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.Status(http.StatusOK)
	}
}

func patchPairingCloseH(a lib.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		err := a.ClosePairingSession(gc.Request.Context())
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.Status(http.StatusOK)
	}
}
