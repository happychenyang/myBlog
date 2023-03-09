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
    }, {
      id: 3,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 4,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 5,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 6,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 7,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 8,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 9,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 10,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 11,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 12,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 13,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 14,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 15,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 16,
      image: '',
      title: 'Title 2',
      content: '这里是内容',
      look: 123
    }, {
      id: 17,
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