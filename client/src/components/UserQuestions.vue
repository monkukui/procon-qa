<template>
  <div class="user-questions">
    <hr>
    <h1>質問一覧</h1>
    
    <div v-for="(value, index) in questions" :key=index>
      <!-- answerd とか, questionedTime とかの命名規則を揃える -->
      <!-- 子コンポーネントには QuestionId だけを渡す-->
      <v-expansion-panels>
        <v-expansion-panel style="margin-top: 1%;">
          <v-expansion-panel-header>

            <router-link class="title" :to="{ name: 'question', query: { questionId: value.id }}" >{{ value.title }}</router-link >
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            回答数: {{ value.answerCount }}
            いいね: {{ value.favoriteCount }}
            {{ value.completed }}
            {{ value.date }}
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
      <!--QuestionPanel
        :questionId="value.id"
      /-->
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

  // FIXME boolean の変数を作る
  private tr: boolean = true;
  private fal: boolean = false;

  private curPageId: number = 1;

  private user: string = '';
  // FIXME any
  private questions = [];

  private created(): void {
    // this.getQuestions();
    this.getQuestionsWithPage();
  }

  // 質問をページ取得する
  private getQuestionsWithPage(): void {
    const url = '/api/user-questions/' + String(this.curPageId);
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
