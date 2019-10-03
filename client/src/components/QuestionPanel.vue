<template>
  <div class="question-panel">
    <v-card
      max-width="1000"
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
  
      <v-row style="margin-left: 1%;">
        <v-col md="1">
          <SquarePanel 
            message="回答数"
            num=3
            color="blue-grey"
          />
        </v-col>
        <v-col md="1">
          <SquarePanel 
            message="閲覧数"
            num=10
            color="blue-grey"
          />
        </v-col>
        <v-col md="1">
          <SquarePanel
            v-if="question.completed"
            message="解決済"
            color="#5cb85c"
          />
          <SquarePanel
            v-if="!question.completed"
            message="未解決"
            color="#f0ad4e"
          />
        </v-col>
        <div class="flex-grow-1"></div>
        <div style="margin-left:auto; margin-right: 5%; margin-top: 5%;">
          <span>{{ question.date }} {{ userName }}</span>
        </div>
      </v-row>
    </v-card>
    
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';

@Component({
  components: {
    SquarePanel,
  },
})
export default class QuestionPanel extends Vue {
  @Prop()
  private questionId!: number;

  // データベース Question 通り
  private question = {
    title: 'title title title title title',
    body: 'body body body body body',
    completed: false,
  };

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
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.question = json;
      this.setUser();
    });
  }

  private setUser(): void {

    const url = 'api/user/' + String(this.question.uid);
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
