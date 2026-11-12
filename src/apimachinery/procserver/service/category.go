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

package service

import (
	"context"
	"net/http"

	"configcenter/src/common/errors"
	"configcenter/src/common/metadata"
)

// CreateServiceCategory TODO
func (s *service) CreateServiceCategory(ctx context.Context, h http.Header,
	data map[string]interface{}) (resp *metadata.Response, err error) {
	resp = new(metadata.Response)
	subPath := "/create/proc/service_category"

	err = s.client.Post().
		WithContext(ctx).
		Body(data).
		SubResourcef(subPath).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

// DeleteServiceCategory TODO
func (s *service) DeleteServiceCategory(ctx context.Context, h http.Header,
	data map[string]interface{}) (resp *metadata.Response, err error) {
	resp = new(metadata.Response)
	subPath := "/delete/proc/service_category"

	err = s.client.Delete().
		WithContext(ctx).
		Body(data).
		SubResourcef(subPath).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

// SearchServiceCategory TODO
func (s *service) SearchServiceCategory(ctx context.Context, h http.Header, opt *metadata.ListServiceCategoryOption) (
	*metadata.MultipleServiceCategory, errors.CCErrorCoder) {

	resp := new(metadata.MultipleServiceCategoryResult)
	subPath := "/findmany/proc/service_category"

	err := s.client.Post().
		WithContext(ctx).
		Body(opt).
		SubResourcef(subPath).
		WithHeaders(h).
		Do().
		Into(resp)

	if err != nil {
		return nil, errors.CCHttpError
	}
	if err := resp.CCError(); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// UpdateServiceCategory TODO
func (s *service) UpdateServiceCategory(ctx context.Context, h http.Header,
	data map[string]interface{}) (resp *metadata.Response, err error) {
	resp = new(metadata.Response)
	subPath := "/update/proc/service_category"

	err = s.client.Put().
		WithContext(ctx).
		Body(data).
		SubResourcef(subPath).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}
