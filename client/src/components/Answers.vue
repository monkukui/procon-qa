<template>
  <div class="answers">
    <div v-for="(value, index) in answers" :key=index>
      <AnswerPanelDetail
        :answerId="value.id"
      />
    </div>
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

  private questionId: number;
  private answers = [];

  private created(): void {
    this.questionId = Number(this.$route.query.questionId);
    this.getAnswers();
  }

  // 質問をページ取得する
  private getAnswers(): void {
    const url = 'api/answers/' + String(this.questionId);
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

}
</script>

<style scoped>
.answers {
  margin: 0px 0px 20px-24px;
}
</style>
