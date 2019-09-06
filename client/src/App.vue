<template>
<!-- App.vue -->

<v-app>
  <v-navigation-drawer app :dark="tr">
    <!-- -->
  </v-navigation-drawer>

  <v-app-bar app :absolute="tr">
    <!-- -->
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

  <!-- Sizes your content based upon application components -->
  <v-content>

    <!-- Provides the application the proper gutter -->
    <v-container fluid>

      <!-- If using vue-router -->
      <router-view></router-view>
    </v-container>
  </v-content>

  <v-footer app>
    <!-- -->
  </v-footer>
</v-app>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  data(): function() {
    return {
      tr: true,
    }
  }
});
</script>
<style>
</style>
