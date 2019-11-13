<template>
  <div class="question-panel">
    <v-card
      class="mx-auto"
      color="white"
    >
      <div v-if="isReady">
        <v-card-title>
          <router-link class="title" :to="{ name: 'question', query: { questionId: this.questionId }}" >{{ question.title }}</router-link >
        </v-card-title>
        <!--v-btn
          class="tag"
          color="blue-grey lighten-4"
          x-small
        >
          タグ {{ question.tid }}
        </v-btn-->
        <v-divider class="mx-4"></v-divider>
        <v-row style="margin-left: 1%;">
          <v-col col="12" sm="1">
            <SquarePanel 
              message="回答数"
              :num="question.answerCount"
              color="blue-grey"
            />
          </v-col>
          <v-col col="12" sm="1" style="margin-left: 2%">
            <SquarePanel
              message="閲覧数"
              num=10
              color="blue-grey"
            />
          </v-col>
          <v-col col="12" sm="1" style="margin-left: 2%">
            <SquarePanel
              message="いいね"
              :num="question.favoriteCount"
              color="blue-grey"
            />
          </v-col>
          <v-col col="12" sm="1" style="margin-left: 2%">
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
          <div style="margin-left:auto; margin-right: 5%; margin-top: 10%;">
            <span class="date">{{ question.date }} {{ userName }}</span>
          </div>
        </v-row>
      </div>
      <div v-else>
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
import { Component, Prop, Vue, Emit, Watch } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';
import {Question} from '@/models/Question.ts';

@Component({
  components: {
    SquarePanel,
  },
})
export default class QuestionPanel extends Vue {
  @Prop({required: true})
  private questionId!: number;

  // データベース Question 通り
  private question: Question = {
    id: 0,
    title: 'title title title title title',
    body: 'body body body body body',
    completed: false,
    date: '',
  };

  private isReady: boolean = false;

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
    const url = '/api/no-auth/question/' + String(this.questionId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.question = json;
      this.setUser();
      this.isReady = true;
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
* {
  color: rgb(66,66,66);
}
.date {
  color: rgb(151, 151, 151);
}
</style>
