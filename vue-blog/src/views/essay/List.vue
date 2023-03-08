<template>
  <div class="div-list">
    <EssayList :pagination="pagination" :data="data"></EssayList>
  </div>
</template>

<style scoped>
.div-list{
  width: 60%;
  margin: 1% auto auto;
}
</style>
<script>
import {defineComponent} from "vue";
import EssayList from "@/components/EssayList";

export default defineComponent({
  name: "list",
  components: {
    EssayList
  },
  setup() {
    const pagination = {
      onChange: page => {
        console.log(page);
      },
      pageSize: 10
    };
    const data = [{
      image: '/image/wood.jpg',
      title: 'Title 1',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }];
    return {
      pagination,
      data
    };
  }
});
</script>