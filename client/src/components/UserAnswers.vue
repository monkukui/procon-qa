<template>
  <div class="user-answers">
    <hr>
    <h1>回答一覧</h1>
    
    <div v-for="(value, index) in answers" :key=index>
      <!-- answerd とか, answeredTime とかの命名規則を揃える -->
      <!-- 子コンポーネントには QuestionId だけを渡す-->
      <AnswerPanelDetail
        :answerId="value.id"
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
import AnswerPanelDetail from '@/components/AnswerPanelDetail.vue';
@Component({
  components: {
    AnswerPanelDetail,
  },
})
export default class Answers extends Vue {

  private curPageId: number = 1;

  private user: string = '';
  // FIXME any
  private answers = [];

  private created(): void {
    // this.getQuestions();
    this.getAnswersWithPage();
  }

  // 質問をページ取得する
  private getAnswersWithPage(): void {
    const url = '/api/user-answers/' + String(this.curPageId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.answers = [];
      this.answers = json;
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
