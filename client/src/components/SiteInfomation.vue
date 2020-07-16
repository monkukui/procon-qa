<template>
  <div class="site-information">
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
        <v-toolbar-title class="title">サイト情報</v-toolbar-title>
      </v-app-bar>
      <div v-if="
        isReadyTotalUser &&
        isReadyTotalAnswer && 
        isReadyCompletedRate && 
        isReadyTotalQuestion &&
        isReadyTotalAnswerGood && 
        isReadyTotalQuestionGood
      ">
        <v-list-item three-line>
          <v-list-item-content>
            <br>
            <v-list dense>
              <v-list-item v-if="isReadyTotalQuestion">
                <v-list-item-content>
                  <v-list-item-title>質問数</v-list-item-title>
                </v-list-item-content>
                <v-list-item-content>
                  <v-list-item-title>{{ totalQuestions }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-list-item v-if="isReadyTotalAnswer">
                <v-list-item-content>
                  <v-list-item-title>回答数</v-list-item-title>
                </v-list-item-content>
                <v-list-item-content>
                  <v-list-item-title>{{ totalAnswers }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-list-item v-if="isReadyTotalUser">
                <v-list-item-content>
                  <v-list-item-title>ユーザ数</v-list-item-title>
                </v-list-item-content>
                <v-list-item-content>
                  <v-list-item-title>{{ totalUsers }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-list-item v-if="isReadyCompletedRate">
                <v-list-item-content>
                  <v-list-item-title>解決済み</v-list-item-title>
                </v-list-item-content>
                <v-list-item-content>
                  <v-list-item-title>{{ completedRate }} %</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-list-item-content>
        </v-list-item>
      </div>
      <div v-else>
        <v-skeleton-loader
          class="mx-auto"
          max-width="344"
          type="list-item"
        ></v-skeleton-loader>
        <v-skeleton-loader
          class="mx-auto"
          max-width="344"
          type="list-item"
        ></v-skeleton-loader>
        <v-skeleton-loader
          class="mx-auto"
          max-width="344"
          type="list-item"
        ></v-skeleton-loader>
        <v-skeleton-loader
          class="mx-auto"
          max-width="344"
          type="list-item"
        ></v-skeleton-loader>
      </div>
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
  private totalQuestionGood: number = 0;
  private totalAnswerGood: number = 0;
  private completedRate: string = '0.0';

  private isReadyTotalQuestion: boolean = false;
  private isReadyCompletedRate: boolean = false;
  private isReadyTotalAnswer: boolean = false;
  private isReadyTotalUser: boolean = false;
  private isReadyTotalQuestionGood: boolean = false;
  private isReadyTotalAnswerGood: boolean = false;

  private created(): void {
    this.getTotalQuestion();
    this.getTotalAnswer();
    this.getTotalUser();
    this.getTotalQuestionGood();
    this.getTotalAnswerGood();
  }

  // 質問へのいいね数を取得する
  private getTotalQuestionGood(): void {
    const url = '/api/no-auth/question-good/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalQuestionGood = cnt;
      this.isReadyTotalQuestionGood = true;
    });
  }

  // 回答へのいいね数を取得する
  private getTotalAnswerGood(): void {
    const url = '/api/no-auth/answer-good/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalAnswerGood = cnt;
      this.isReadyTotalAnswerGood = true;
    });
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
      this.isReadyTotalQuestion = true;
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
      this.completedRate = (cnt * 100 / this.totalQuestions).toFixed(3);
      this.isReadyCompletedRate = true;
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
      this.isReadyTotalAnswer = true;
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
      this.isReadyTotalUser = true;
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
