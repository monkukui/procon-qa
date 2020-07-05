<template>
  <div class="answers">

    <div v-if="isReady">
      <div v-for="(value, index) in answers" :key=index>
        <AnswerPanelDetail
          :answerId="value.id"
        />
      </div>
    </div>
    <div v-else class="text-center">  
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
import AnswerPanelDetail from '@/components/AnswerPanelDetail.vue';
@Component({
  components: {
    AnswerPanelDetail,
  },
})
export default class Answers extends Vue {
  @Prop()
  private mode!: number;
  @Prop()
  private xor!: boolean;
  private questionId!: number;
  private answers = [];
  private isReady: boolean = false;

  private created(): void {
    this.questionId = Number(this.$route.query.questionId);
    this.getAnswers();
  }

  // 回答を全取得
  private getAnswers(): void {
    this.isReady = false;
    const url = 'api/no-auth/answers/' + String(this.questionId) + '/' + String(this.mode);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.answers = [];
      this.answers = json;
      this.isReady = true;
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }
  // mode が変われば，質問を再取得
  @Watch('mode')
  private modeChanged(): void {
    this.getAnswers();
  }

  @Watch('xor')
  private xorChanged(): void {
    this.getAnswers();
  }

}
</script>

<style scoped>
.answers {
  margin: 10px;
}
</style>
