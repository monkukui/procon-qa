<template>
  <div class="question-panel-detail">
    <v-card
      max-width="1010"
      class="mx-auto"
      color="white"
    >
      <v-card-title>
        {{ question.title }}
      </v-card-title>
      <v-btn
        class="tag"
        color="blue-grey lighten-4"
        x-small
      >
        tag (仮)
      </v-btn>
      
      <v-divider class="mx-4"></v-divider>
      <v-card-text>{{ question.body }}</v-card-text>
      <v-card-text>質問日時: {{ question.date }}</v-card-text>
      <v-card-text>url: {{ question.url }}</v-card-text>
      <v-card-text>ステータス: {{ question.state }}</v-card-text>
      <v-card-text>回答状況: {{ question.completed }}</v-card-text>
    </v-card>
    
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class QuestionPanel extends Vue {
  // 元々 string として良いのでは ??
  private questionId: number;

  // データベース Question 通り
  private question = {};

  private created(): void {
    // TODO api/question/:id (GET) を叩く
    this.questionId = Number(this.$route.query['questionId']);
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
.small {
  font-size: 75%;
}
.question-panel-detail {
  margin: 3%;
}

.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}
</style>
