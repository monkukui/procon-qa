<template>
  <div class="answer-panel-detail">
    <v-card
      max-width="1010"
      class="mx-auto"
      color="white"
    >
      <v-card-title>
        {{ answer.body }}
      </v-card-title>
      
      <v-divider class="mx-4"></v-divider>
      <v-card-text>回答日時: {{ answer.date }}</v-card-text>
      <v-card-text>回答者: {{ userName }}</v-card-text>
      <v-card-text>{{ answer.favo }} ファボ</v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class AnswerPanelDetail extends Vue {
  @Prop()
  private answerId!: number;

  private answer = {};
  /*
    "uid": number,
    "qid": number,
    "id": number,
    "body": string,
    "date": string,
    "favo": number,
  */

  private userName: string = '';

  private created(): void {
    this.createAnswer();
  }


  private createAnswer(): void {
    const url = 'api/answer/' + String(this.answerId);
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

    const url = 'api/user/' + String(this.answer.uid);
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

.small {
  font-size: 75%;
}

.answer-panel-detail {
  margin: 2%;
}
  
</style>
