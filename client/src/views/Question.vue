<template>
  <div class="question">
    <TopInfo />
    <p>個別の質問ページ</p>
    <QuestionPanelDetail
      :title="title"
      :body="body"
      :questionedTime="questionedTime"
    />
    <br>

    <!-- Answers を作る -->
    <Answers />
    <AnswerForm />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TopInfo from '@/components/TopInfo.vue';
import QuestionPanelDetail from '@/components/QuestionPanelDetail.vue';
import AnswerForm from '@/components/AnswerForm.vue';
import Answers from '@/components/Answers.vue';

@Component({
  components: {
    TopInfo,
    QuestionPanelDetail,
    AnswerForm,
    Answers,
  },
})
export default class Question extends Vue {
  // TODO getParam とかでとってくる
  // TODO any は最悪なのでなんとかする (string | undefined とかにしてもエラーがでる)
  private title: any = '';
  private body: any = '';
  private questionedTime: any = '';

  private mounted() {
    this.title = this.$route.query.title;
    this.body = this.$route.query.body;
    this.questionedTime = this.$route.query.questionedTime;
  }
}
</script>
