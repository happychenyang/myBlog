<template>
  <a-list
      :grid="{ gutter: 16,  column: 1  }"
      :data-source="data"
      :pagination="pagination"
  >
    <template #renderItem="{ item }" >
      <router-link to="/detail" :id="item.id">
        <a-list-item>
          <a-card :title="item.title">
            <img  :src="item.image ? item.image : defaultImage" alt="" :width="130" :height="90">
            {{ item.content }}
          </a-card>
        </a-list-item>
      </router-link>
    </template>
  </a-list>
</template>

<script setup>
//
defineProps({
  pagination: Object,
  data: Array
})
const defaultImage = '/image/bear.jpeg'
</script>
