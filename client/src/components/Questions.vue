<template>
  <div class="questions">
    <v-pagination
      v-model="curPageId"
      :length="length"
      :total-visible="7"
    ></v-pagination>
    <div v-if="isReady">
      <div v-for="(value, index) in questions" :key=index>
        <!-- answerd とか, questionedTime とかの命名規則を揃える -->
        <!-- 子コンポーネントには QuestionId だけを渡す-->
        <QuestionPanel
          :questionId="value.id"
        />
      </div>
      <v-pagination
        v-model="curPageId"
        :length="length"
        :total-visible="7"
      ></v-pagination>
    </div>
    <div v-else>
      <v-progress-circular
        :size="100"
        color="primary"
        indeterminate
      ></v-progress-circular>
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
export default class Questions extends Vue {

  @Prop()
  private mode!: number;

  // FIXME boolean の変数を作る
  private tr: boolean = true;
  private fal: boolean = false;

  private curPageId: number = 1;

  private isReady: boolean = false;

  private user: string = '';
  // FIXME any
  private questions = [];
  private totalQuestions: number = 0;
  private length: number = 1;

  private created(): void {
    // this.getQuestions();
    this.getQuestionsWithPage();
    this.getTotalQuestion();
  }

  // 質問数を取得する
  private getTotalQuestion(): void {
    const url = '/api/no-auth/questions/count';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((cnt) => {
      this.totalQuestions = cnt;
      this.length = (this.totalQuestions + 9) / 10; // 切り上げ
    });
  }

  // 質問をページ取得する
  private getQuestionsWithPage(): void {
    this.isReady = false;
    const url = '/api/no-auth/questions/' + String(this.curPageId) + '/' + String(this.mode);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.questions = [];
      this.questions = json;
      this.isReady = true;
    });
  }

  // 質問を全取得する
  private getQuestions(): void {
    const url = '/api/no-auth/questions/' + String(this.curPageId);
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

  // mode が変われば，質問
  @Watch('mode')
  private modeChanged(): void {
    this.getQuestionsWithPage();
  }
}
</script>

<style scoped>
.questions {
}
</style>
