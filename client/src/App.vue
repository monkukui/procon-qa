<template>
<!-- App.vue -->
<v-app>
  <v-navigation-drawer
    app
    dark
    v-model="drawer"
    :clipped="$vuetify.breakpoint.lgAndUp"
  >
    <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            メニュー
          </v-list-item-title>
          <v-list-item-subtitle>
            サブタイトル
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
        dense
        nav
      >
        <v-list-item
          key="ホーム"
          link
          to="/"
        >
          <v-list-item-icon>
            <v-icon>mdi-home</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>ホーム</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-list-item
          key="マイページ"
          link
          to="/mypage"
        >
          <v-list-item-icon>
            <v-icon>mdi-account</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>マイページ</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        
        <v-list-item
          key="タグ検索"
          link
          to="/tagsearch"
        >
          <v-list-item-icon>
            <v-icon>mdi-magnify</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>タグ検索</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        
        <v-list-item
          key="このサイトについて"
          link
          to="/about"
        >
          <v-list-item-icon>
            <v-icon>mdi-information</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>このサイトについて</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        
      </v-list>
    <!-- -->
  </v-navigation-drawer>
  
  <div class="bar">
  <v-app-bar
    app
    :clipped-left="$vuetify.breakpoint.lgAndUp"
  >
    <!-- -->
    <v-toolbar-title class="headline text-uppercase">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <span>PROCON</span>
      <span class="font-weight-light">QA</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn
      color="primary"
      to="questionform"
    >
      <span class="mr-2">質問する</span>
    </v-btn>
    <v-btn
      v-if="user"
      text
      @click="logout"
    >
      <span class="mr-2">ログアウト</span>
    </v-btn>
    <v-btn
      v-if="!user"
      text
      to="login"
    >
      <span class="mr-2">ログイン</span>
    </v-btn>
    <v-btn
      v-if="!user"
      text
      to="signup"
    >
      <span class="mr-2">登録</span>
    </v-btn>
  </v-app-bar>
  </div>

  <!-- Sizes your content based upon application components -->
  <v-content>
    <!-- Provides the application the proper gutter -->
    <v-container fluid>
      <!-- If using vue-router -->
      <router-view></router-view>
    </v-container>
  </v-content>
</v-app>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class App extends Vue {
  private drawer: any = null;
  private user: string = '';
  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.user = claims.name;
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
  private logout(): void {
    localStorage.removeItem('token');
    location.reload();
  }
  private isLoggedIn(): boolean {
    return false;
  }

}
</script>

<style scoped>
.bar {
  z-index: 2147483647;
}
</style>
