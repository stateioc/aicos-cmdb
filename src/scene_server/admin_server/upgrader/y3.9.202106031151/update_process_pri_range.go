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

package y3_9_202106031151

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
	"context"
)

func updatePriorityProperty(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	filter := map[string]interface{}{
		common.BKObjIDField:      common.BKInnerObjIDProc,
		common.BKPropertyIDField: "priority",
		common.BKAppIDField:      0,
	}
	doc := map[string]interface{}{
		"option": map[string]interface{}{
			"min": common.MinProcessPrio,
			"max": common.MaxProcessPrio,
		},
		"placeholder": "批量启动进程依据优先级从小到大排序操作，停止进程按反序操作。优先级数值范围支持输入 [ -100 ~ 10000 ]",
	}
	if err := db.Table(common.BKTableNameObjAttDes).Update(ctx, filter, doc); err != nil {
		blog.Errorf("update failed, filter:%#v, doc:%#v, err:%v", filter, doc, err)
		return err
	}

	return nil
}
