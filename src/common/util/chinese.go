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

package util

import (
	"github.com/mozillazg/go-pinyin"
)

// GetInitials get the initials of the input, if the input is chinese, returns the initials of its pinyin
func GetInitials(input string) string {
	if len(input) == 0 {
		return ""
	}

	for _, initials := range input {
		pinyinResult := pinyin.SinglePinyin(initials, pinyin.NewArgs())
		for _, p := range pinyinResult {
			if len(p) > 0 {
				return string(p[0])
			}
		}

		return string(initials)
	}

	return ""
}
