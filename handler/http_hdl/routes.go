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
	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	_ "github.com/SENERGY-Platform/mgw-auth-service/handler/http_hdl/swagger_docs"
	"github.com/SENERGY-Platform/mgw-auth-service/lib"
	"github.com/SENERGY-Platform/mgw-auth-service/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const healthCheckPath = "/health-check"

var routes = gin_mw.Routes[lib.Api]{
	GetServiceHealthH,
	GetIdentitiesH,
	PostIdentityH,
	GetIdentityH,
	PatchIdentityH,
	DeleteIdentityH,
	PostPairingH,
	PatchPairingOpenH,
	PatchPairingCloseH,
	GetSrvInfoH,
}

// SetRoutes
// @title Auth Service API
// @version 0.2.13
// @description Provides access to mgw-core auth functions.
// @license.name Apache-2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func SetRoutes(e *gin.Engine, a lib.Api) error {
	err := routes.Set(a, e, util.Logger)
	if err != nil {
		return err
	}
	e.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	return nil
}
