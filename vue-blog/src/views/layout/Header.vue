<template>
  <div class="header-div">
    <div class="header-div-one">
      <a-menu v-model:selectedKeys="selectedKeys" mode="horizontal" :theme="theme">
        <a-menu-item key="I" disabled>
          <img alt="Vue logo" style="width: 140px" :src="logoSrc" />
        </a-menu-item>
        <a-menu-item key="one">
          <template #icon>
            <appstore-outlined />
          </template>
          <router-link to="/home">首页</router-link>
        </a-menu-item>
        <a-menu-item key="two">
          <template #icon>
            <tag-outlined />
          </template>
          <router-link to="/essay_list">文章</router-link>
        </a-menu-item>
        <a-menu-item key="sheet">
          <template #icon>
            <heart-outlined />
          </template>
          ChatGPT
        </a-menu-item>
        <a-menu-item key="flor">
          <template #icon>
            <smile-outlined />
          </template>
          简介
        </a-menu-item>
      </a-menu>
      <a-switch
          :checked="theme === 'dark'"
          checked-children="Happy"
          un-checked-children="Luck"
          @change="changeTheme"
      />
    </div>
  </div>
</template>
<script>
import { defineComponent, ref, toRefs , reactive } from 'vue';
import { MailOutlined, AppstoreOutlined, SettingOutlined, HeartOutlined, SmileOutlined, TagOutlined } from '@ant-design/icons-vue';
export default defineComponent({
  components: {
    MailOutlined,
    AppstoreOutlined,
    SettingOutlined,
    HeartOutlined,
    SmileOutlined,
    TagOutlined,
  },
  setup() {
    const state = reactive({
      theme: "light",
      selectedKeys: ["one"],
      logoSrc: "/image/blog.jpeg",
    });
    const changeTheme = checked => {
      state.theme = checked ? 'dark' : 'light';
    };
    return {
      ...toRefs(state),
      changeTheme,
    };
  },
});
</script>
<style scoped>
.header-div{
  width: 100%;
  margin: auto;
  background: white;
  position: fixed;
  box-shadow: 0px 5px 5px 0px #D5D5D5;
  z-index: 9999;
}
.header-div > .header-div-one{
  width: 60%;
  margin: auto;
}
.header-div > .header-div-one > .ant-switch{
  float: right;
  margin-top: -45px;
  margin-right: 15px;
}
</style>
