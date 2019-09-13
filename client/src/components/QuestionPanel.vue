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
        タグ {{ question.tid }}
      </v-btn>
      
      <v-divider class="mx-4"></v-divider>
      <v-card-text>質問日時: {{ question.date }}</v-card-text>
      <v-card-text>質問者: {{ userName }}</v-card-text>
    </v-card>

  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch } from 'vue-property-decorator';

@Component
export default class QuestionPanel extends Vue {
  @Prop()
  private questionId!: number;
  
  // データベース Question 通り
  private question = {};
  
  // 質問者の名前
  private userName: string = '';

  // コンポーネントが作られた時に呼び出される関数
  private created(): void {
    this.init(); 
  }

  // questionId が変更されたらば,表示を変える
  @Watch('questionId')
  private idChanged(): void {
    this.init();
  }

  private init(): void {
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
      this.setUser();
    })
  }

  private setUser(): void {
    
    const url = 'api/user/' + String(this.question.uid);
    const headers = {'Authorization': `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then(response => {
      if(response.ok) {
        return response.json();
      }
      return []
    }).then(json => {
      console.log(this.question);
      console.log(this.userName);
      this.userName = json.name;
      console.log(this.userName);
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
