<template>
  <div class="user-answers">
    <hr>
    <h1>回答一覧</h1>
    <div v-for="(value, index) in questions" :key=index>
      <!-- answerd とか, answeredTime とかの命名規則を揃える -->
      <!-- 子コンポーネントには QuestionId だけを渡す-->
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
export default class UserAnswers extends Vue {

  private userId: string | Array<(string | null)> = '';
  private curPageId: number = 1;

  private user: string = '';
  // FIXME any
  private questions = [];

  private created(): void {
    // this.getQuestions();
    this.userId = this.$route.query.uid;
    this.getAnswersWithPage();
  }

  // 質問をページ取得する
  private getAnswersWithPage(): void {
    const url = '/api/no-auth/user-answers/' + this.userId + '/' + String(this.curPageId);
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
    this.getAnswersWithPage();
  }
}
</script>

<style scoped>
.answers {
}
</style>
