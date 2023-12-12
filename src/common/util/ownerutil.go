/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"configcenter/src/common"
	"strconv"
	"strings"
)

// SetQueryOwner returns condition that in default ownerID and request ownerID
func SetQueryOwner(condition map[string]interface{}, ownerID string) map[string]interface{} {
	if nil == condition {
		condition = make(map[string]interface{})
	}
	if ownerID == common.BKSuperOwnerID {
		return condition
	}
	if ownerID == common.BKDefaultOwnerID {
		condition[common.BKOwnerIDField] = common.BKDefaultOwnerID
		return condition
	}
	condition[common.BKOwnerIDField] = map[string]interface{}{common.BKDBIN: []string{common.BKDefaultOwnerID, ownerID}}
	return condition
}

// SetModOwner set condition equal owner id, the condition must be a map or struct
func SetModOwner(condition map[string]interface{}, ownerID string) map[string]interface{} {
	if nil == condition {
		condition = make(map[string]interface{})
	}
	if ownerID == common.BKSuperOwnerID {
		return condition
	}
	condition[common.BKOwnerIDField] = ownerID
	return condition
}

func GenerateKey(fieldName string, index int) string {
	if index == 0 {
		return fieldName
	}
	return fieldName + "-" + strconv.Itoa(index)
}

func GetOriginalKey(fieldName string) string {
	parts := strings.Split(fieldName, "-")
	return parts[0]
}

func ConvertToAndMap(data map[string]interface{}) map[string]interface{} {
	andOrKeys := []string{common.BKDBAND, common.BKDBOR}

	for _, key := range andOrKeys {
		if _, ok := data[key].([]interface{}); ok {
			return data
		}
	}
	andMap := make(map[string]interface{})
	andMap[common.BKDBAND] = []interface{}{}

	for key, value := range data {
		originalKey := GetOriginalKey(key)

		condition := map[string]interface{}{
			originalKey: value,
		}

		andMap[common.BKDBAND] = append(andMap[common.BKDBAND].([]interface{}), condition)
	}

	return andMap
}
