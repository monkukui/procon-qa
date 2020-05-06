<template>
  <div>
    <h1>for debug</h1>
    <h1>ブラウザバックしてください</h1>

    <h1>質問</h1>
    <p>質問数 {{ questionCount }}</p>
    <p>回答数 {{ answerCount }}</p>
    <p>ユーザ数 {{ userCount }}</p>
    <p>質問へのいいね {{ questionGoodCount }}</p>
    <p>回答へのいいね {{ answerGoodCount }}</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';


@Component({
  components: {
  },
})
export default class DataBase extends Vue {
  private questionCount: number = -1000000;
  private answerCount: number = -1000000;
  private userCount: number = -1000000;
  private questionGoodCount: number = -1000000;
  private answerGoodCount: number = -1000000;

  private created(): void {
    this.getQuestionSize();
    this.getAnswerSize();
    this.getUserSize();
    this.getQuestionGoodSize();
    this.getAnswerGoodSize();
  }

  // 質問数取得
  private getQuestionSize(): void {
    const url = '/api/no-auth/questions/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.questionCount = json;
    });
  }
  // 回答数取得
  private getAnswerSize(): void {
    const url = '/api/no-auth/answers/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.answerCount = json;
    });
  }
  // ユーザ数取得
  private getUserSize(): void {
    const url = '/api/no-auth/users/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userCount = json;
    });
  }

  // 質問へのいいね数
  private getQuestionGoodSize(): void {
    const url = '/api/no-auth/question-good/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.questionGoodCount = json;
    });
  }

  // 回答へのいいね数
  private getAnswerGoodSize(): void {
    const url = '/api/no-auth/answer-good/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.answerGoodCount = json;
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }

}
</script>

<style scoped>
</style>
