#!/bin/sh
#######################################
# SPACE-AICOS is pleased to support the open source community by making
# SPACE-AICOS available.
# Copyright (C) 2026 China Academy of Information and Communications Technology (CAICT).
# All rights reserved.
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://opensource.org/licenses/MIT
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on
# an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
# either express or implied. See the License for the
# specific language governing permissions and limitations under the License.
#######################################

set -e

PACKAGES=$(go list ../src/...)

#test:
	echo "go test"
	go test -cover=true ../src/...

#collect-cover-data:
	echo "collect-cover-data"
	echo "mode: count" > coverage-all.out
	for pkg in $PACKAGES;do
		echo "collect package:"${pkg}
		go test -v -coverprofile=coverage.out -covermode=count ${pkg};\
		if [ -f coverage.out ]; then\
			tail -n +2 coverage.out >> coverage-all.out;\
		fi\
	done

#test-cover-html:
	echo "test-cover-html"
	go tool cover -html=coverage-all.out -o coverage.html

#test-cover-func:
	echo "test-cover-func"
	go tool cover -func=coverage-all.out

#get total result
	total=`go tool cover -func=coverage-all.out | tail -n 1 | awk '{print $3}'`
	echo "total coverage: "${total}
	

