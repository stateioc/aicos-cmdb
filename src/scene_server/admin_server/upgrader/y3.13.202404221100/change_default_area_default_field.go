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

package y3_13_202404221100

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func changeDefaultAreaDefaultField(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	cond := mapstr.MapStr{
		common.BKCloudIDField: common.BKDefaultDirSubArea,
		common.BKDefaultField: mapstr.MapStr{common.BKDBNE: int64(common.BuiltIn)},
	}
	data := mapstr.MapStr{common.BKDefaultField: int64(common.BuiltIn)}

	if err := db.Table(common.BKTableNameBasePlat).Update(ctx, cond, data); err != nil {
		blog.Errorf("update cloud area failed, cond: %+v, data: %+v, err: %v", cond, data, err)
		return err
	}

	return nil
}
