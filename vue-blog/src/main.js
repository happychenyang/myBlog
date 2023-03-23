import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
import "github-markdown-css"
import config from '@/request/config';

const appObj = createApp(App);
appObj.config.globalProperties.httpConfig = config
appObj.use(store).use(router).use(Antd).mount("#app");
