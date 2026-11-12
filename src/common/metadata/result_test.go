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

package metadata

import (
	"testing"

	"configcenter/src/common/errors"
	"configcenter/src/common/metadata"
)

func TestResponse(t *testing.T) {

	err := errors.New(9999999, "test-msg")

	respPtr := &metadata.Response{
		BaseResp: metadata.BaseResp{
			Result: false,
			Code:   err.GetCode(),
			ErrMsg: err.Error(),
		},
	}

	ccErr := respPtr.CCError()
	if ccErr == nil {
		t.Errorf("not error")
		return
	}

	if err.GetCode() != ccErr.GetCode() ||
		err.Error() != ccErr.Error() {
		t.Errorf("code info, code:%v, error msg:%s", ccErr.GetCode(), ccErr.Error())
		return
	}

	resp := metadata.Response{
		BaseResp: metadata.BaseResp{
			Result: false,
			Code:   err.GetCode(),
			ErrMsg: err.Error(),
		},
	}

	ccErr = resp.CCError()
	if ccErr == nil {
		t.Errorf("not error")
		return
	}

	if err.GetCode() != ccErr.GetCode() ||
		err.Error() != ccErr.Error() {
		t.Errorf("code info, code:%v, error msg:%s", ccErr.GetCode(), ccErr.Error())
		return
	}

	respSucc := metadata.Response{
		BaseResp: metadata.BaseResp{
			Result: true,
			Code:   0,
			ErrMsg: "",
		},
	}
	ccErr = respSucc.CCError()
	if ccErr != nil {
		t.Errorf("have error")
		return
	}

}
