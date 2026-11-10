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

<template functional>
  <i class="loading" v-if="props.loading"></i>
  <div v-else><slot></slot></div>
</template>

<style lang="scss" scoped>
  .loading {
    display: inline-block;
    width: 16px;
    height: 16px;
    background-color: transparent;
    background-image: url("../../assets/images/icon/loading.svg");
    background-position: center center;
    background-size: 16px;
    background-repeat: no-repeat;
  }
</style>
