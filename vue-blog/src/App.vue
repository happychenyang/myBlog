<template>
  <a-layout>
    <a-layout-header :style="{ padding: '0px 0px' }">
      <Header />
    </a-layout-header>
    <a-layout-content>
      <router-view></router-view>
    </a-layout-content>
    <a-layout-footer>
      <Footer />
    </a-layout-footer>
  </a-layout>
  <div>
    <meta name="viewport" content="width=device-width, initial-scale=1">
  </div>
</template>
<script>
import {defineComponent} from 'vue';
import Header from "@/views/layout/Header";
import Footer from "@/views/layout/Footer";

export default defineComponent({
  components: {
    Header,
    Footer
  },
});
</script>
<style>
.ant-layout > .ant-layout-header {
  background: #fff;
}
</style>