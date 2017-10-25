// Copyright 2013 wetalk authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"github.com/astaxie/beego/orm"

	"github.com/cuu/select_tags/models"
)

type ModelsGetSearchController struct {
	BaseController
}


// @router /model/pick [post]
func (this *ModelsGetSearchController) ModelPickPost() {
	id := this.GetString("id")
	model := this.GetString("model")
	
	result := map[string]interface{}{
		"success": false,
	}

	var data []orm.ParamsList

	defer func() {
		if len(data) > 0 {
			result["success"] = true
			result["data"] = data[0]
		}
		this.Data["json"] = result
		this.ServeJSON()
	}()

	var qs orm.QuerySeter

	switch model {
	case "User":
		qs = models.Users()
	case "Nutrition":
		qs = models.Nutritions()
	case "Ingredient":
		qs = models.Ingredients()
	case "Dish":
		qs = models.Dishes()
	}

	qs = qs.Filter("Id", id).Limit(1)

	switch model {
	case "User":
		qs.ValuesList(&data, "Id", "UserName")
	case "Nutrition":
		qs.ValuesList(&data,"Id","Name")
	case "Ingredient":
		qs.ValuesList(&data,"Id","Name")
	case "Dish":
		qs.ValuesList(&data,"Id","Name")
	}
	
}


// @router /model/select [post]
func (this *ModelsGetSearchController) ModelSelectPost() {
	search := this.GetString("search")
	model := this.GetString("model")
	result := map[string]interface{}{
		"success": false,
	}

	var data []orm.ParamsList

	defer func() {
		if len(data) > 0 {
			result["success"] = true
			result["data"] = data
		}
		this.Data["json"] = result
		this.ServeJSON()
	}()

	if len(search) < 3 {
		return
	}

	switch model {
	case "User":
		models.Users().Filter("UserName__icontains", search).Limit(10).ValuesList(&data, "Id", "UserName")

	case "Nutrition":
		models.Nutritions().Filter("Name__icontains", search).Limit(10).ValuesList(&data, "Id", "Name")
	case "Ingredient":
		models.Ingredients().Filter("Name__icontains", search).Limit(10).ValuesList(&data,"Id","Name")
	case "Dish":
		models.Dishes().Filter("Name__icontains",search).Limit(10).ValuesList(&data,"Id","Name")
	}
}
