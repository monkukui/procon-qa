<template>
<!-- App.vue -->
  <v-app>
    <div class="bar">
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
              :to="{ name: 'userpage', query: { uid: this.uid }}"
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
      </div>

      <div class="bar">
      <v-app-bar
        app
        color="white"
        height="80"
        :clipped-left="$vuetify.breakpoint.lgAndUp"
      >
        <!-- -->
        <v-toolbar-title class="headline text-uppercase">
          <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
          <router-link to="/">
            <img class="logo" src="./assets/square_logo_white.jpg" alt="logo" width="70" height="auto">
          </router-link>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn
          v-if="user"
          text
          :to="{ name: 'userpage', query: { uid: this.uid }}"
        >
          {{ user }}
        </v-btn>
        <v-btn
          v-if="user"
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
    <div class="footer">
      <v-footer
        absolute
        class="font-weight-medium"
      >
        <v-col
          class="text-center"
          cols="12"
        >
          {{ new Date().getFullYear() }} — <strong>procon-qa</strong>
        </v-col>
      </v-footer>
    </div>
  </v-app>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class App extends Vue {
  private drawer: boolean = true;
  private user: string = '';
  private uid: string = '';
  private closeModal: boolean = false;
  private created(): void {
    // 認証が必要な api を叩いてみて，その結果によって分岐
    const url = '/api/token';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        const token = this.getToken();
        if (token !== null) {
          const claims = JSON.parse(atob(this.getToken().split('.')[1]));
          this.user = claims.name;
          this.uid = claims.uid;
        }
      } else {
        localStorage.removeItem('token');
      }
    });
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
.draw {
  z-index: 2147483647 - 1;
}
.logo {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
}
.footer {
  margin-bottom: 100px;
}
</style>
