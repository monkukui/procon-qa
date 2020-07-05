<template>
<!-- App.vue -->
  <v-app>
    <div class="bar">
      <v-navigation-drawer
        dark
        app
        v-model="drawer"
        clipped="$vuetify.breakpoint.lgAndUp"
      >
          <v-list
            dense
            nav
          >
            <v-list-item
              key="ホーム"
              link
              to="/"
              @click="clickLink"
            >
              <v-list-item-icon>
                <v-icon>mdi-home</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>ホーム</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item
              key="通知"
              link
              to="/notification"
              @click="clickLink"
            >
              <v-list-item-icon>
                <v-icon>mdi-bell</v-icon>
                <v-icon v-if="notificationFlag" color="rgb(73, 160, 237)" class="notification_draw" size="large">mdi-checkbox-blank-circle</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>通知</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item
              key="マイページ"
              link
              :to="{ name: 'userpage', query: { uid: this.uid }}"
              @click="clickLink"
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
              @click="clickLink"
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
              @click="clickLink"
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
          <v-app-bar-nav-icon @click.stop="drawer = !drawer" @click="clickLink"></v-app-bar-nav-icon>
          <v-icon v-if="notificationFlag" color="rgb(73, 160, 237)" class="notification_bar" size="large">mdi-checkbox-blank-circle</v-icon>
          <router-link 
            to="/"
          >
            <img @click="clickLink" class="logo" src="./assets/square_logo_white.jpg" alt="logo" width="70" height="auto">
          </router-link>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn
          v-if="user"
          text
          :to="{ name: 'userpage', query: { uid: this.uid }}"
          @click="clickLink"
        >
          <div style="text-transform: lowercase;">
            {{ user }}
          </div>
        </v-btn>
        <v-btn
          v-if="user"
          color="primary"
          to="questionform"
          @click="clickLink"
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
          <v-btn icon href="https://github.com/monkukui/procon-qa">
            <v-icon color="rgb(37, 41, 46)">mdi-github</v-icon>
          </v-btn>
          <v-btn icon href="https://twitter.com/monkukui2">
            <v-icon color="rgb(73, 160, 236)">mdi-twitter</v-icon>
          </v-btn>
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
  private drawer: boolean = false;
  private user: string = '';
  private uid: string = '';
  private closeModal: boolean = false;
  private notificationFlag: boolean = false; // 通知があるかどうかのフラグ

  private created(): void {
    // 認証が必要な api を叩いてみて，その結果によって分岐
    const url = '/api/token';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        const token = this.getToken();
        if (token !== null) {
          const claims = JSON.parse(atob(token.split('.')[1]));
          this.uid = claims.uid;
        }

        this.setUser();

      } else {
        localStorage.removeItem('token');
      }
    });
  }

  private clickLink(): void {
    this.setUser();
  }
  private setUser(): void {
    if (this.uid === '') {
      return;
    }


    const urlForUserName = '/api/no-auth/user/' + String(this.uid);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(urlForUserName, {headers}).then((res) => {
      if (res.ok) {
        return res.json();
      }
      return [];
    }).then((json) => {
      this.user = json.name;
      this.notificationFlag = json.notification_flag;
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }
  private logout(): void {
    localStorage.removeItem('token');
    location.reload();
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
.notification_draw {
  position: relative;
  left: -12px;
  top: -8px;
}
.notification_bar {
  position: relative;
  left: -20px;
  top: -10px;
}
</style>
