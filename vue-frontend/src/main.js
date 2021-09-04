import Vue from 'vue'
import App from './App.vue'
import router from "./router/router"

import "./assets/css/normalize.css";

import 'element-ui/lib/theme-chalk/index.css';
import ElementUI from "element-ui";
import lang from "element-ui/lib/locale/lang/ru-RU";
import locale from "element-ui/lib/locale";
locale.use(lang)
Vue.use(ElementUI, { locale });


Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')