<template>
  <div id="app">
    <Header />
    <router-view />
  </div>
</template>

<script>
import Vue from "vue";

import axios from "axios";
import VueAxios from "vue-axios";

import Header from "@/components/layout/Header.vue";

export default {
  name: "App",
  components: { Header },
  created() {
    const hasura = axios.create({
      baseURL: "http://195.133.48.245:8081/v1/graphql",
      withCredentials: true,
      headers: { "x-hasura-admin-secret": "qwerty" },
    });

    const fastApi = axios.create({
      baseURL: "http://195.133.48.245:10199/",
    });

    Vue.use(VueAxios, { hasura, fastApi });
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
