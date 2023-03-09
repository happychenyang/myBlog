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
        <div v-for=" item in imageList ">
          <img :src="item" alt="" style="width: 100%; height: 300px; margin: auto"/>
        </div>
      </a-carousel>
    </div>
    <div class="content-div-content">
      <EssayList :pagination=" pagination " :data=" data "></EssayList>
    </div>
  </div>
</template>
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
<script setup>
import { defineComponent, ref } from "vue";
import { StarOutlined, LikeOutlined, MessageOutlined, UserOutlined } from "@ant-design/icons-vue";
import EssayList from "@/components/EssayList";

const pagination = {
  onChange: page => {
    console.log(page);
  },
  pageSize: 5
};
const imageList = [
  '/image/wood.jpg',
  '/image/rainbow.jpg',
  '/image/waterfall.jpg',
  '/image/sunset.jpg'
];
const value = ref();
const onPanelChange = (value, mode) => {
  console.log(value, mode);
};
const data = [{
  id: 1,
  image: '/image/wood.jpg',
  title: 'Title 1',
  content: '这里是内容',
  look: 123
}, {
  id: 2,
  image: '',
  title: 'Title 2',
  content: '这里是内容',
  look: 123
}];
</script>