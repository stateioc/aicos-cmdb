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

import { ref, watch } from 'vue'
import { t } from '@/i18n'
import { $bkInfo } from '@/magicbox/index.js'
import { clone } from '@/utils/tools'

export default function useSideslider(data, options = {}) {
  const isChanged = ref(false)
  const { watchOnce = true } = options
  // 默认弹出框文本信息
  const defaultInfoData = {
    subTitle: '离开将会导致未保存信息丢失',
    title: '确认离开当前页？',
    okText: '离开',
    cancelText: '取消'
  }
  let infoData = clone(defaultInfoData)

  if (data) {
    // 放到下次任务循环队列执行，因为枚举多选类型一开始为空值，后面第一次正常赋值这块会执行
    setTimeout(() => {
      const unwatch = watch(data, () => {
        isChanged.value = true
        watchOnce && unwatch()
      }, { deep: true })
    }, 300)
  }

  const beforeClose = (confirmCallback, cancelCallback) => new Promise((resolve, reject) => {
    if (!isChanged.value) {
      confirmCallback && confirmCallback?.()
      resolve(true)
      return
    }
    const { subTitle, title, okText, cancelText } = infoData
    $bkInfo({
      title: t(title),
      subTitle: t(subTitle),
      clsName: 'custom-info-confirm default-info',
      okText: t(okText),
      cancelText: t(cancelText),
      confirmFn() {
        confirmCallback && confirmCallback?.()
        resolve(true)
      },
      cancelFn() {
        cancelCallback && cancelCallback?.()
        reject(false)
      },
    })
  })

  const reset = () => {
    setTimeout(() => {
      isChanged.value = false
    })
  }

  const setChanged = (v) => {
    isChanged.value = v
  }

  const setInfoData = (data = {}) => {
    infoData = Object.assign({}, defaultInfoData, data)
  }
  return {
    beforeClose,
    isChanged,
    reset,
    setChanged,
    setInfoData
  }
}
