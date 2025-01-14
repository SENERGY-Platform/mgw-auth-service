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
	"time"
)

type pairingOpenQuery struct {
	Duration int64 `form:"duration"`
}

// PostPairingH godoc
// @Summary Pair device
// @Description Transmit device information to pair a device.
// @Tags Device Pairing
// @Accept json
// @Produce	json
// @Param meta body map[string]any true "device information"
// @Success	200 {object} lib_model.CredentialsResponse "generated credentials"
// @Failure	400 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /pairing/request [post]
func PostPairingH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, path.Join(lib_model.PairingPath, lib_model.PairingReqPath), func(gc *gin.Context) {
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

// PatchPairingOpenH godoc
// @Summary Start paring
// @Description Open paring endpoint and create a paring session.
// @Tags Device Pairing
// @Param duration query integer false "set session duration in nanoseconds (default=5m)"
// @Success	200
// @Failure	400 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /pairing/open [patch]
func PatchPairingOpenH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPatch, path.Join(lib_model.PairingPath, lib_model.OpenPath), func(gc *gin.Context) {
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

// PatchPairingCloseH godoc
// @Summary Stop pairing
// @Description Close paring endpoint and cancel paring session.
// @Tags Device Pairing
// @Success	200
// @Failure	500 {string} string "error message"
// @Router /pairing/close [patch]
func PatchPairingCloseH(a lib.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPatch, path.Join(lib_model.PairingPath, lib_model.ClosePath), func(gc *gin.Context) {
		err := a.ClosePairingSession(gc.Request.Context())
		if err != nil {
			_ = gc.Error(err)
			return
		}
		gc.Status(http.StatusOK)
	}
}
