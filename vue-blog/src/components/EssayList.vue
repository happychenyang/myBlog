<template>
  <a-list
      :grid="{ gutter: 16,  column: 1  }"
      :data-source="data"
      :pagination="pagination"
  >
    <template #renderItem="{ item }" >
      <a-list-item>
        <a-card :title="item.title">
          <img  :src="item.image ? item.image : defaultImage" alt="" :width="130" :height="90">
          {{ item.content }}
        </a-card>
      </a-list-item>
    </template>
  </a-list>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  name: "EssayList",
  props:{
    pagination: Object,
    data: Array
  },
  setup() {
    const defaultImage = '/image/bear.jpeg'
    return {
      defaultImage
    };
  },
});
</script>
