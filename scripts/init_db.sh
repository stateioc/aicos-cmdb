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

#!/bin/bash
set -e

# get local IP.
localIp=`python ip.py`

# 判断是否为IPV6，是则在地址两端加中括号
if [[ ${localIp} =~ ":" ]]
then
  localIp="[${localIp}]"
fi
echo "localIp:${localIp}"

curl -X POST -H 'Content-Type:application/json' -H 'X-Bkcmdb-User:migrate' -H 'X-Bkcmdb-Supplier-Account:0' http://${localIp}:60004/migrate/v3/migrate/community/0

echo ""
