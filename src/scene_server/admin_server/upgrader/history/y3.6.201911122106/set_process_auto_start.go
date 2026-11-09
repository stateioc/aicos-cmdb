/*
 * SPACE-AICOS is pleased to support the open source community by making
 * SPACE-AICOS available.
 * Copyright (C) 2026 China Academy of Information and Communications Technology (CAICT).
 * All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package y3_6_201911122106

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

// setProcessAutoStartAttribute TODO
// set process model's attribute auto_start's bk_isapi field value to true
func setProcessAutoStartAttribute(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	cond := map[string]string{
		"bk_obj_id":      "process",
		"bk_property_id": "auto_start",
	}
	target := map[string]bool{
		"bk_isapi": true,
	}
	if err := db.Table(common.BKTableNameObjAttDes).Update(ctx, cond, target); err != nil {
		return err
	}

	return nil
}
