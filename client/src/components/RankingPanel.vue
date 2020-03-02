<template>
  <div class="ranking-panel">
    <v-card
      class="mx-auto"
      max-width="344"
      outlined
    >
      <v-app-bar
        color="rgb(66, 66, 66)"
        flat
        dense
      >
        <v-toolbar-title class="title">ランキング</v-toolbar-title>
      </v-app-bar>
      <v-list-item three-line>
        <v-list-item-content>
          <br>
          <v-list dense>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>順位</v-list-item-title>
              </v-list-item-content>
              <v-list-item-content>
                <v-list-item-title>ユーザ</v-list-item-title>
              </v-list-item-content>
              <v-list-item-content>
                <v-list-item-title>獲得いいね数</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-for="(user, i) in users"
              :key="i"
            >
              <v-list-item-content>
                {{ i + 1 }}
              </v-list-item-content>
              <v-list-item-content>
                <!--UserName
                  :name="user.name"
                  :color="user.color"
                /-->
                <UserName
                  :name="user.name"
                />
              </v-list-item-content>
              <v-list-item-content>
                <v-list-item-title>{{ user.favorite_sum }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>

        </v-list-item-content>
      </v-list-item>
    </v-card>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';
import UserName from '@/components/UserName.vue';

@Component({
  components: {
    UserName,
  },
})
export default class RankingPanel extends Vue {
  private users: any = {};
  private created(): void {
    // user をページ取得する
    const url = '/api/no-auth/users/1/1';
    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.users = json;
    });
  }
}
</script>

<style>
.title {
  color: rgb(256, 256, 256);
}
</style>
