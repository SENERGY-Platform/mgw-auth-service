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

package main

import (
	"context"
	"errors"
	"fmt"
	sb_logger "github.com/SENERGY-Platform/go-service-base/logger"
	"github.com/SENERGY-Platform/mgw-auth-service/handler/cs_hdl"
	"github.com/SENERGY-Platform/mgw-auth-service/handler/http_hdl"
	"github.com/SENERGY-Platform/mgw-auth-service/handler/kratos_hdl"
	lib_model "github.com/SENERGY-Platform/mgw-auth-service/lib/model"
	"github.com/SENERGY-Platform/mgw-auth-service/service"
	"github.com/SENERGY-Platform/mgw-auth-service/util"
	srv_info_hdl "github.com/SENERGY-Platform/mgw-go-service-base/srv-info-hdl"
	sb_util "github.com/SENERGY-Platform/mgw-go-service-base/util"
	"github.com/SENERGY-Platform/mgw-go-service-base/watchdog"
	kratos "github.com/ory/kratos-client-go"
	"net"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"
)

var version string

func main() {
	srvInfoHdl := srv_info_hdl.New("auth-service", version)

	ec := 0
	defer func() {
		os.Exit(ec)
	}()

	util.ParseFlags()

	config, err := util.NewConfig(util.Flags.ConfPath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		ec = 1
		return
	}

	logFile, err := util.InitLogger(config.Logger)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		var logFileError *sb_logger.LogFileError
		if errors.As(err, &logFileError) {
			ec = 1
			return
		}
	}
	if logFile != nil {
		defer logFile.Close()
	}

	util.Logger.Printf("%s %s", srvInfoHdl.GetName(), srvInfoHdl.GetVersion())

	util.Logger.Debugf("config: %s", sb_util.ToJsonStr(config))

	watchdog.Logger = util.Logger
	wtchdg := watchdog.New(syscall.SIGINT, syscall.SIGTERM)

	kratosConf := kratos.NewConfiguration()
	kratosConf.Servers = []kratos.ServerConfiguration{
		{
			URL: config.HttpClient.IdentitySrvBaseUrl,
		},
	}
	kratosClient := kratos.NewAPIClient(kratosConf)
	identityHdl := kratos_hdl.New(kratosClient, time.Duration(config.HttpClient.Timeout))

	csHdl := cs_hdl.New(time.Duration(config.CSDefDuration))
	defer csHdl.Close()

	srv := service.New(identityHdl, csHdl, srvInfoHdl)

	httpHandler, err := http_hdl.New(srv, map[string]string{
		lib_model.HeaderApiVer:  srvInfoHdl.GetVersion(),
		lib_model.HeaderSrvName: srvInfoHdl.GetName(),
	})
	if err != nil {
		util.Logger.Error(err)
		ec = 1
		return
	}

	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(config.ServerPort), 10))
	if err != nil {
		util.Logger.Error(err)
		ec = 1
		return
	}
	server := &http.Server{Handler: httpHandler}
	srvCtx, srvCF := context.WithCancel(context.Background())
	wtchdg.RegisterStopFunc(func() error {
		if srvCtx.Err() == nil {
			ctxWt, cf := context.WithTimeout(context.Background(), time.Second*5)
			defer cf()
			if err := server.Shutdown(ctxWt); err != nil {
				return err
			}
			util.Logger.Info("http server shutdown complete")
		}
		return nil
	})
	wtchdg.RegisterHealthFunc(func() bool {
		if srvCtx.Err() == nil {
			return true
		}
		util.Logger.Error("http server closed unexpectedly")
		return false
	})

	wtchdg.Start()

	diCtx, diCF := context.WithCancel(context.Background())
	wtchdg.RegisterStopFunc(func() error {
		diCF()
		return nil
	})

	srv.CreateInitialIdentity(diCtx, config.InitIdentity.User, config.InitIdentity.Secret.Value(), time.Second*5, 10)

	go func() {
		defer srvCF()
		util.Logger.Info("starting http server ...")
		if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
			util.Logger.Error(err)
			ec = 1
			return
		}
	}()

	ec = wtchdg.Join()
}
