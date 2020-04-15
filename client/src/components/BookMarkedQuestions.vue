<template>
  <div class="user-questions">
    <h1>ブックマークした質問一覧</h1>
    <h2>{{ totalQuestions }} 件の質問</h2>
    <v-pagination
      v-model="curPageId"
      :length="length"
      :total-visible="7"
    ></v-pagination>
    <div v-for="(value, index) in questions" :key=index>
      <QuestionPanel
        :questionId="value.id"
      />
    </div>
    
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch} from 'vue-property-decorator';
import QuestionPanel from '@/components/QuestionPanel.vue';

@Component({
  components: {
    QuestionPanel,
  },
})
export default class BookMarkedQuestions extends Vue {

  private userId: string | Array<(string | null)> = '';
  // FIXME boolean の変数を作る
  private tr: boolean = true;
  private fal: boolean = false;

  private curPageId: number = 1;

  private user: string = '';
  // FIXME any
  private questions = [];
  private totalQuestions: number = 0;
  private length: number = 1;

  private created(): void {
    // this.getQuestions();
    this.userId = this.$route.query.uid;
    this.getQuestionsWithPage();
    this.getTotalQuestion();
  }

  // 質問数を取得する
  private getTotalQuestion(): void {
    const url = '/api/no-auth/book-marked-questions/count/' + this.userId;
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalQuestions = cnt;
      this.length = Math.ceil(this.totalQuestions / 10); // 切り上げ
    });
  }

  // 質問をページ取得する
  private getQuestionsWithPage(): void {
    const url = '/api/no-auth/book-marked-questions/' + this.userId + '/' + String(this.curPageId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.questions = [];
      this.questions = json;
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }

  // ページが変われば, 質問を取得し直す
  @Watch('curPageId')
  private pageChanged(): void {
    this.getQuestionsWithPage();
  }
}
</script>

<style scoped>
.title {
  color: #0288D1;
  text-decoration: none;
}
.title:hover {
  color: #29B6F6;
  text-decoration: none;
}
</style>
