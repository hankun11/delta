/* Copyright (C) 2017 Beijing Didi Infinity Technology and Development Co.,Ltd.
All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
==============================================================================*/
package handler

import (
	"delta/deltann/server/core/conf"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

type ModelVersion struct {
	Version string      `json:"version"`
	State   string      `json:"state"`
	Status  ModelStatus `json:"status"`
}

type ModelStatus struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func DeltaModelHandler(context *gin.Context) {
	defer glog.Flush()
	modelVersion := &ModelVersion{conf.DeltaConf.Model.Graph[0].Version, "AVAILABLE", ModelStatus{"OK", ""}}
	pagesJson, err := json.Marshal(modelVersion)
	if err != nil {
		glog.Infof("Cannot encode to JSON %s ", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"model_version_status": string(pagesJson)})
}
