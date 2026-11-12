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

package params

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpecialCharChange(t *testing.T) {
	type testUnit struct {
		src string
		dst string
	}

	testUnits := []testUnit{
		{
			src: ".",
			dst: `\.`,
		},
		{
			src: "(",
			dst: `\(`,
		},
		{
			src: ")",
			dst: `\)`,
		},
		{
			src: "\\",
			dst: `\\`,
		},
		{
			src: "|",
			dst: `\|`,
		},
		{
			src: "[",
			dst: `\[`,
		},
		{
			src: "]",
			dst: `\]`,
		},
		{
			src: "-",
			dst: `\-`,
		},
		{
			src: "*",
			dst: `\*`,
		},
		{
			src: "{",
			dst: `\{`,
		},
		{
			src: "}",
			dst: `\}`,
		},
		{
			src: "^",
			dst: `\^`,
		},
		{
			src: "$",
			dst: `\$`,
		},
		{
			src: "?",
			dst: `\?`,
		},

		{
			src: "aaa",
			dst: `aaa`,
		},
		{
			src: "12345",
			dst: `12345`,
		},
		{
			src: "12345",
			dst: `12345`,
		},
		{
			src: "!@#%&_+=,<>/`~",
			dst: "!@#%&_+=,<>/`~",
		},
		{
			src: "12345676890qwertyuiopasdfghjklzxcvbnm QWERTYUIOPASDFGHJKLZXCVBNM",
			dst: "12345676890qwertyuiopasdfghjklzxcvbnm QWERTYUIOPASDFGHJKLZXCVBNM",
		},
	}

	for _, item := range testUnits {
		dst := SpecialCharChange(item.src)
		// blog.Infof("src:%s, expact dst:%s, dst:%s", item.src, item.dst, dst)
		require.Equal(t, item.dst, dst)
	}

}
