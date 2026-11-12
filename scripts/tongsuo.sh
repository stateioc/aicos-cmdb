#!/bin/bash
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

if [ "${DISABLE_CRYPTO}" = true ];then
  echo "crypto is disabled"
  exit 0
fi

if [ -d "$TONGSUO_PATH" ]; then
  echo "tongsuo already exists"
  exit 0
fi

echo -e "\033[34mpreparing tongsuo... \033[0m"

wget --no-check-certificate https://github.com/Tongsuo-Project/Tongsuo/archive/refs/tags/8.3.2.tar.gz
tar zxvf 8.3.2.tar.gz
cd Tongsuo-8.3.2/

if [ "$IS_STATIC" == true ]; then
  yum -y install glibc-static
  ./config --prefix=$TONGSUO_PATH -static
else
  ./config --prefix=$TONGSUO_PATH
fi

make -j
make install
