<template>
  <a-list
      :grid="{ gutter: 16,  column: 1  }"
      :data-source="data"
      :pagination="pagination"
  >
    <template #renderItem="{ item }" >
      <router-link :to=" { path:'/detail', query:{ id: item.id }} ">
        <a-list-item>
          <a-card :title="item.name">
            <img  :src=" global.httpConfig.baseUrl + (item.article_image ? item.article_image : defaultImage)" alt="" :width="130" :height="90">
            {{ item.introduction }}
          </a-card>
        </a-list-item>
      </router-link>
    </template>
  </a-list>
</template>

<script setup>
import { ref, watch, getCurrentInstance } from 'vue'

//获取全局上下文参数
const { appContext } = getCurrentInstance()
const global = appContext.config.globalProperties;

//传参数接受
const props = defineProps({
  pagination: Object,
  data: Array
})

const defaultImage = '/image/bear.jpeg'
</script>
