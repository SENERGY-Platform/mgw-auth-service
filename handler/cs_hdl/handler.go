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

package cs_hdl

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	lib_model "github.com/SENERGY-Platform/mgw-auth-service/lib/model"
	"github.com/SENERGY-Platform/mgw-auth-service/util"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Handler struct {
	cChan       chan struct{}
	open        bool
	sID         string
	defDuration time.Duration
	mu          sync.Mutex
}

func New(defaultDuration time.Duration) *Handler {
	return &Handler{
		defDuration: defaultDuration,
		cChan:       make(chan struct{}, 1),
	}
}

func (h *Handler) Open(duration time.Duration) error {
	if duration == 0 {
		duration = h.defDuration
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if !h.open {
		id, err := uuid.NewUUID()
		if err != nil {
			return lib_model.NewInternalError(fmt.Errorf("generating session ID failed: %s", err))
		}
		h.sID = id.String()
		util.Logger.Infof("credential session '%s' open", h.sID)
		h.open = true
		go h.run(duration)
	} else {
		return lib_model.NewInternalError(fmt.Errorf("credential session '%s' already open", h.sID))
	}
	return nil
}

func (h *Handler) Close() {
	h.mu.Lock()
	if h.open {
		h.open = false
		h.cChan <- struct{}{}
	}
	h.mu.Unlock()
}

func (h *Handler) GetCredentials() (string, string, error) {
	h.mu.Lock()
	defer func() {
		h.mu.Unlock()
	}()
	if h.open {
		var err error
		defer func() {
			if err == nil {
				h.open = false
				h.cChan <- struct{}{}
			}
		}()
		var s string
		s, err = getRandom()
		if err != nil {
			return "", "", lib_model.NewInternalError(fmt.Errorf("credential session '%s': generating login failed: %s", h.sID, err))
		}
		l := s
		s, err = getRandom()
		if err != nil {
			return "", "", lib_model.NewInternalError(fmt.Errorf("credential session '%s': generating secret failed: %s", h.sID, err))
		}
		return l, s, nil
	}
	return "", "", errors.New("no credential session open")
}

func (h *Handler) run(duration time.Duration) {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-timer.C:
		util.Logger.Debugf("credential session '%s' timed out", h.sID)
		break
	case <-h.cChan:
		break
	}
	h.mu.Lock()
	h.open = false
	util.Logger.Infof("credential session '%s' closed", h.sID)
	h.sID = ""
	h.mu.Unlock()
}

func getRandom() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write([]byte(id.String()))
	return hex.EncodeToString(hash.Sum(nil)), nil
}
