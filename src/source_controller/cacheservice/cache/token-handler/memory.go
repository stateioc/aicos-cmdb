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

package tokenhandler

import (
	"context"

	"configcenter/src/storage/stream/types"
)

var _ types.TokenHandler = new(MemoryHandler)

// MemoryHandler is a token handler that stores the token in process memory
type MemoryHandler struct {
	token string
}

// NewMemoryTokenHandler generate a new memory event token handler
func NewMemoryTokenHandler() *MemoryHandler {
	return new(MemoryHandler)
}

// SetLastWatchToken set last event watch token
func (m *MemoryHandler) SetLastWatchToken(ctx context.Context, token string) error {
	m.token = token
	return nil
}

// GetStartWatchToken get event start watch token
func (m *MemoryHandler) GetStartWatchToken(ctx context.Context) (string, error) {
	return m.token, nil
}
