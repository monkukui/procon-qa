<template>
<!-- App.vue -->
<v-app>
  <v-navigation-drawer app :dark="tr">
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

  <v-app-bar app>
    <!-- -->
    <v-toolbar-title class="headline text-uppercase">
      <span>PROCON</span>
      <span class="font-weight-light">QA {{ user }}</span>
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
     text
     to="signup"
    >
      <span class="mr-2">登録</span>
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
    <v-col
      class="text-center"
      cols="12"
    >
      {{ new Date().getFullYear() }} — <strong>PROCONQA</strong>
    </v-col>
  </v-footer>
</v-app>
</template>

<script lang="ts">

import { Component, Prop, Vue, Emit, Watch} from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class App extends Vue {

  private tr: boolean = true;
  private fal: boolean = false;

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
  }
  private isLoggedIn(): boolean {
    return false;
  }

}
</script>

</script>
<style>
</style>
