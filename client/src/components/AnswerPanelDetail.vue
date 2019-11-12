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
    
    <v-card-actions>
      <v-card-text>投稿日時: {{ userName }}</v-card-text>
      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-share-variant</v-icon>
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
  /*
    "uid": number,
    "qid": number,
    "id": number,
    "body": string,
    "date": string,
    "favo": number,
  */

  // どうにかならんか?
  private tr: boolean = true;
  private fa: boolean = false;

  private userName: string = '';

  private created(): void {
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
