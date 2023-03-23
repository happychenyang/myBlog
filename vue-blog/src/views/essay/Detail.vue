<template>
  <div class="div-detail">
    <h1>{{ articleObj.title }}</h1>
    <div v-html="articleObj.content" id="content"></div>
  </div>
</template>

<script setup>

import {onMounted, reactive, ref, watch} from "vue";
import {marked} from "marked"
import API from "@/request/api"
import { useRoute } from "vue-router";

const route = useRoute();


// marked 选项
marked.setOptions({
  pedantic: false,
  gfm: true,
  breaks: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
  xhtml: false,
});

let articleObj = reactive({
  title: '',
  content: ''
})

onMounted (() => {
  const id = Number(route.query.id)
  API.article.articleDetail(id).then(res => {
    articleObj.content = marked(res.Data.ArticleText.article_content)
    articleObj.title = res.Data.name
  });
});

</script>

<style scoped>
.div-detail {
  width: 60%;
  margin: 1% auto;
}
</style>