<!--
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
-->

<script setup>
  import Notice from '@blueking/notice-component-vue2'
  import '@blueking/notice-component-vue2/dist/style.css'
  import { ref, onMounted, onBeforeUnmount } from 'vue'

  const emit = defineEmits(['show-change', 'size-change'])

  const dataUrl = `${window.API_HOST}notice/get_current_announcements`
  const noticeComp = ref(null)

  const noticeShowChange = (isShow) => {
    emit('show-change', isShow)
  }

  const execResizeHander = () => {
    const el = noticeComp.value.$el
    const height = el.offsetHeight
    const width = el.offsetHeight
    emit('size-change', [width, height])
  }
  const resizeObserver = new ResizeObserver((entries) => {
    for (const entry of entries) {
      if (entry.target === noticeComp.value.$el) {
        execResizeHander()
      }
    }
  })
  onMounted(() => {
    resizeObserver.observe(noticeComp.value.$el)
  })
  onBeforeUnmount(() => {
    resizeObserver.unobserve(noticeComp.value.$el)
  })
</script>

<template>
  <notice ref="noticeComp" :api-url="dataUrl" @show-alert-change="noticeShowChange" />
</template>
