<template>
  <div class="site-information">
    <v-card
      class="mx-auto"
      max-width="344"
      outlined
    >
      <v-app-bar
        color="primary"
        flat
        dense
      >
        <v-toolbar-title class="title">サイト情報</v-toolbar-title>
      </v-app-bar>
      <v-list-item three-line>
        <v-list-item-content>
          <br>
          <v-list-item>質問数:  {{ totalQuestions }}</v-list-item>
          <v-list-item>回答数:  {{ totalAnswers }}</v-list-item>
          <v-list-item>ユーザ数:  {{ totalUsers }}</v-list-item>
          <v-list-item>解決済み:  {{ completedRate }} %</v-list-item>
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
  private totalAnswers: number = 0;
  private totalUsers: number = 0;
  private completedRate: string = '0.0';

  private created(): void {
    this.getTotalQuestion();
    this.getTotalAnswer();
    this.getTotalUser();
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
      this.getCompletedRate();
    });
  }
  // 解決済み率数を取得する
  private getCompletedRate(): void {
    const url = '/api/no-auth/completed-questions/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.completedRate = (cnt / this.totalQuestions).toFixed(3);
    });
  }
  // 回答数を取得する
  private getTotalAnswer(): void {
    const url = '/api/no-auth/answers/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalAnswers = cnt;
    });
  }

  // ユーザ数を取得する
  private getTotalUser(): void {
    const url = '/api/no-auth/users/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalUsers = cnt;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style>
.title {
  color: rgb(256, 256, 256);
}
</style>
