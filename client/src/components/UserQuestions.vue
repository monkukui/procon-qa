<template>
  <div class="user-questions">
    <hr>
    <h1>質問一覧</h1>
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
export default class UserQuestions extends Vue {

  private userId: string | Array<(string | null)> = '';
  // FIXME boolean の変数を作る
  private tr: boolean = true;
  private fal: boolean = false;

  private curPageId: number = 1;

  private user: string = '';
  // FIXME any
  private questions = [];

  private created(): void {
    // this.getQuestions();
    this.userId = this.$route.query.uid;
    this.getQuestionsWithPage();
  }

  // 質問をページ取得する
  private getQuestionsWithPage(): void {
    const url = '/api/no-auth/user-questions/' + this.userId + '/' + String(this.curPageId);
    alert(url);
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

  // 質問を全取得する
  private getQuestions(): void {
    const url = '/api/questions/' + String(this.curPageId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
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
