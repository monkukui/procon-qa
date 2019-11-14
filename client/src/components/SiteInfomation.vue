<template>
  <div class="site-information">
    <v-card
    class="mx-auto"
    max-width="344"
    shaped
  >
    <v-list-item three-line>
      <v-list-item-content>
        <v-list-item-title class="headline">サイト情報</v-list-item-title>
        <br>
        <v-list-item>質問数:  {{ totalQuestions }}</v-list-item>
        <v-list-item>回答数:  2321</v-list-item>
        <v-list-item>ユーザ数:  4320</v-list-item>
        <v-list-item>回答率:  64%</v-list-item>
      </v-list-item-content>
    </v-list-item>
  </v-card>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class SiteInformation extends Vue {

  private totalQuestions: number = 0;
  private created(): void {
    this.getTotalQuestion();
  }

  // 質問数を取得する
  private getTotalQuestion(): void {
    const url = '/api/no-auth/questions/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalQuestions = cnt;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style>
</style>
