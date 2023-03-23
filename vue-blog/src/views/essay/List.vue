<template>
  <div class="div-list">
    <EssayList :pagination=" pagination " :data=" data.list "></EssayList>
  </div>
</template>

<style scoped>
.div-list {
  width: 60%;
  margin: 1% auto auto;
}
</style>
<script setup>
import {defineComponent, reactive} from "vue";
import EssayList from "@/components/EssayList";
import API from "@/request/api"

const pagination = {
  onChange: page => {
    pagination.pageIndex = page
    getList()
  },
  pageIndex: 1,
  pageSize: 10
};

// 声明这是个一个动态
let data = reactive({list: []})

const getList = () => {
  API.article.articleList(pagination).then(res => {
    data.list = res.Data
    pagination.total = res.Total
  });
}

getList(1);

</script>