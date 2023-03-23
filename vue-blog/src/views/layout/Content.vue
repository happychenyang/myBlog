<template>
  <div class="content-div">
    <div class="content-div-left">
      <a-card hoverable>
        <template #cover>
          <img alt="example" src="/image/bear.jpeg"/>
        </template>
        <a-card-meta title="个人留言">
          <template #description>本人很懒，没留下什么想说的</template>
        </a-card-meta>
      </a-card>
      <a-card hoverable>
        <a-calendar v-model:value="value" :fullscreen="false" @panelChange="onPanelChange"/>
      </a-card>
    </div>
    <div class="content-div-title">
      <a-carousel autoplay>
        <div v-for=" item in data.imageList ">
          <img :src="global.httpConfig.baseUrl  + item.blog_image" alt="" style="width: 100%; height: 300px; margin: auto"/>
        </div>
      </a-carousel>
    </div>
    <div class="content-div-content">
      <EssayList :pagination=" pagination " :data=" data.list "></EssayList>
    </div>
  </div>
</template>

<script setup>
import {ref, watch, reactive, onMounted, getCurrentInstance} from "vue";
import EssayList from "@/components/EssayList";
import API from "@/request/api"

const { appContext } = getCurrentInstance()
const global = appContext.config.globalProperties;

const pagination = {
  onChange: page => {
    pagination.pageIndex = page;
    getList()
  },
  pageIndex:1,
  pageSize: 10,
  total: 0
};
const imageList = [
  '/image/wood.jpg',
  '/image/rainbow.jpg',
  '/image/waterfall.jpg',
  '/image/sunset.jpg'
];
const value = ref();
const onPanelChange = (value, mode) => {
  // console.log(value, mode);
};
// 声明这是个一个动态
let data = reactive({
  list: [],
  imageList: []
})

const getList = () => {
  API.article.articleList(pagination).then(res => {
    data.list = res.Data
    pagination.total = res.Total
  });
}

getList(1);

onMounted(() => {
  API.home.rotationList().then(res => {
    data.imageList = res.Data
  });
})

</script>

<style scoped>
.ant-card {
  margin: auto;
  width: 100%;
  box-shadow: none;
}

/* For demo */
.ant-carousel {
  width: 100%;
  margin: auto;
}

.ant-carousel :deep(.slick-slide) {
  text-align: center;
  height: 300px;
  line-height: 300px;
  background: #364d79;
  overflow: hidden;
}

.ant-carousel :deep(.slick-slide h3) {
  color: #fff;
}

.content-div {
  width: 61%;
  height: 100%;
  margin: 1% auto auto;
}

.content-div-title {
  width: 69%;
  height: 300px;
  float: right;
}

.content-div-left {
  width: 30%;
  height: 100%;
  float: left;
}

.content-div-content {
  width: 69%;
  float: right;
  margin: 1% auto auto;
}
</style>