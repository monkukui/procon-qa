<template>
  <div class="answer-panel-detail">
    <v-card
      class="mx-auto"
    >
    <v-row>
      <v-col md="8"> 
        <div class="mavon-editor">
          <mavon-editor
            v-model="answer.body"
            defaultOpen="preview"
            :toolbarsFlag="fa"
            :subfield="fa"
            :boxShadow="fa"
          />
        </div>
      </v-col>
      <v-col md="4">
        <v-card
          class="mx-auto"
          max-width="300"
          outlined
          color="deep-orange lighten-5"
        >
          <v-list-item three-line>
            <v-list-item-content>
              <div>回答者</div>
              <v-list-item-title class="headline mb-1">{{ userName }}</v-list-item-title>
            </v-list-item-content>
            <v-list-item-avatar
              circle
              size="80"
              color="grey"
            ></v-list-item-avatar>
          </v-list-item>
        </v-card>
      </v-col>
    </v-row>
  
    <v-chip
      class="ma-2"
      color="pink"
      text-color="white"
      :disabled="userName == name"
      @click="favoriteAnswer"
    >
      <v-avatar
        left
        class="pink darken-4"
      >
      {{ answer.favoriteCount }}
      </v-avatar>
      いいね
    </v-chip>
    <v-card-actions>
      <v-card-text>投稿日時: {{ userName }}</v-card-text>
      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-share-variant</v-icon>
      </v-btn> 
      <v-btn icon>
        <v-icon @click="deleteAnswer">mdi-home</v-icon>
      </v-btn> 
    </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';
import {Answer} from '@/models/Answer.ts';

@Component
export default class AnswerPanelDetail extends Vue {
  @Prop()
  private answerId!: number;

  private answer: Answer = {
    id: 0,
    uid: 0,
    qid: 0,

    body: '',
    date: '',
  };

  // どうにかならんか?
  private tr: boolean = true;
  private fa: boolean = false;

  // 順に，回答者の名前，ユーザの名前
  private userName: string = '';
  private name: string = '';

  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.name = claims.name;
    this.createAnswer();
  }

  private createAnswer(): void {
    const url = 'api/no-auth/answer/' + String(this.answerId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.answer = json;
      this.setUser();
    });
  }

  private setUser(): void {

    const url = 'api/no-auth/user/' + String(this.answer.uid);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userName = json.name;
    });
  }

  private deleteAnswer(): void {
    const url = 'api/answer/' + String(this.answerId);
    const method = 'DELETE';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        alert('回答を削除しました');
        location.reload();
        // リフレッシュ
      }
    });
  }

  private favoriteAnswer(): void {
    const url = 'api/answer/' + String(this.answerId) + '/favorite';
    const method = 'put';
    const headers = {authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      this.createAnswer();
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style scoped>
.answer-panel-detail {
  margin-top: 5px;
}
.mavon-editor {
  z-index: 0;
}
</style>
