<template>
  <div class="question-panel">
    <v-card
      max-width="1010"
      class="mx-auto"
      color="white"
    >
      <v-card-title>
        <router-link class="title" :to="{ name: 'question', query: { questionId: this.questionId }}" >{{ question.title }}</router-link >
      </v-card-title>
      <v-btn
        class="tag"
        color="blue-grey lighten-4"
        x-small
      >
        タグ(仮)
      </v-btn>
      
      <v-divider class="mx-4"></v-divider>
      <v-card-text>質問日時: {{ question.date }}</v-card-text>
    </v-card>

  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class QuestionPanel extends Vue {
  @Prop()
  private questionId!: number;
  
  // データベース Question 通り
  private question = {};

  private created(): void {
    // TODO api/question/:id (GET) を叩く
    const url = 'api/question/' + String(this.questionId);
    const headers = {'Authorization': `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then(response => {
      if(response.ok) {
        return response.json();
      }
      return []
    }).then(json => {
      this.question = json;
    })
  }
  
  private getToken(): any {
    return localStorage.getItem('token');
  }

}
</script>

<style scoped>
.question-panel {
  margin: 3%;
}

.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}

.title {
  color: #0288D1;
  text-decoration: none;
}
.title:hover {
  color: #29B6F6;
  text-decoration: none;
}

.small {
  font-size: 75%;
}

.question-panel {
}
  
</style>
