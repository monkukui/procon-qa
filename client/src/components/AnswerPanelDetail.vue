<template>
  <div class="answer-panel-detail">
    <v-card
      class="mx-auto"
    >
      <div v-if="isReady">
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
          <v-btn icon
            color="pink"
            :disabled="name == userName || name == ''"
          >
            <v-icon @click="favoriteAnswer">mdi-heart</v-icon>
          </v-btn>
          <v-btn icon>
            <v-icon>mdi-share-variant</v-icon>
          </v-btn> 
          <v-btn icon
            color="error"
            :disabled="name != userName"
          >
            <v-icon @click="alert = !alert">mdi-delete</v-icon>
          </v-btn> 
        </v-card-actions>
        <v-row>
          <v-col md="12">
            <v-alert
              outlined
              :value="alert"
              transition="scale-transition"
            >
              <v-col col="12" sm="8">
                <h1>本当に削除しますか?</h1>
                <p>この操作は取り消せません.回答に付与された「いいね」も取り消されます.</p>
              </v-col>
              <v-col col="12" sm="4">
                <v-btn color="error" @click="deleteAnswer">削除する</v-btn>
                &nbsp;
                <v-btn @click="alert = !alert">戻る</v-btn>
              </v-col>
            </v-alert>
          </v-col>
        </v-row>
      </div>
      <div v-else class="text-center">
        <v-progress-circular
          :size="100"
          color="primary"
          indeterminate
        ></v-progress-circular>
      </div>
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

  // ロード中の制御
  private isReady: boolean = false;

  // 削除の確認画面の制御
  private alert: boolean = false;

  private created(): void {
    if (this.getToken() != null) {
      const claims = JSON.parse(atob(this.getToken().split('.')[1]));
      this.name = claims.name;
    }
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
      this.isReady = true;
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
