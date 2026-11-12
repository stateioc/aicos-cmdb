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

pushd ${BASH_SOURCE%/*} >/dev/null

# list the current directories
DIRS=$(find * -maxdepth 0 -type d | grep cmdb_)

# into the directory to stop the all cmdb process
for tmp_dir in $DIRS;do
    pushd $(pwd) > /dev/null
        num=`ps -efww | grep $tmp_dir | grep -v grep | grep -v tail | wc -l`
        if [ "$num" -gt 0 ];then
            if [ -f "$tmp_dir/stop.sh" ];then
                cd $tmp_dir && bash stop.sh || true
            fi
        fi
    popd > /dev/null
done

for tmp_dir in $DIRS;do
    num=`ps -efww | grep $tmp_dir | grep -v grep | grep -v tail | wc -l`
    if [ "$num" -ge 1 ];then
        echo "Stoped: $tmp_dir"
    fi
done

ps -ef | grep [c]mdb_ | grep -v tail || true
cnt=$(pgrep cmdb_ | grep -v tail | wc -l)
echo "Running process count: $cnt"

popd >/dev/null
