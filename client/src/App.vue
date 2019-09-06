<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>PROCON</span>
        <span class="font-weight-light">QA</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        text
        href="#"
      >
        <span class="mr-2">Login</span>
      </v-btn>
      <v-btn
        text
        href="#"
      >
        <span class="mr-2">SignUp</span>
      </v-btn>
    </v-app-bar>
    
    <v-content class="content">
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import HelloWorld from './components/HelloWorld.vue';
import Home from './views/Home.vue';


export default Vue.extend({
  name: 'App',
  components: {
    HelloWorld,
    Home,
  },
  data: () => ({
    //
  }),
});
</script>
<style>
.content {
  margin: 20px;
}
</style>
