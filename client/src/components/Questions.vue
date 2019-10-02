<template>
  <div class="questions">
    
    <!--div v-for="question in questions">
      {{ question }}
    </div-->
    <v-pagination
      v-model="curPageId"
      :length="20"
      :total-visible="7"
      circle
    ></v-pagination>

    <div v-for="(value, index) in questions" :key=index>
      <!-- answerd とか, questionedTime とかの命名規則を揃える -->
      <!-- 子コンポーネントには QuestionId だけを渡す-->
      <QuestionPanel
        :questionId="value.id"
      />
    </div>

    <v-pagination
      v-model="curPageId"
      :length="20"
      :total-visible="7"
      circle
    ></v-pagination>

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

  // FIXME boolean の変数を作る
  private tr: boolean = true;
  private fal: boolean = false;

  private curPageId: number = 1;

  private user: string = '';
  // FIXME any
  private questions = [];
  private changePageId(id: number): void {
    this.curPageId = id;
  }

  private created(): void {
    // this.getQuestions();
    this.getQuestionsWithPage();
  }

  // 質問をページ取得する
  private getQuestionsWithPage(): void {
    const url = '/api/questions/' + String(this.curPageId);
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
.questions {
}
</style>
