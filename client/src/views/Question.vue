<template>
  <div class="question">
    <QuestionPanelDetail
      :title="title"
      :body="body"
      :questionedTime="questionedTime"
    />

    <Answers />
    <AnswerForm />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QuestionPanelDetail from '@/components/QuestionPanelDetail.vue';
import Answers from '@/components/Answers.vue';

@Component({
  components: {
    QuestionPanelDetail,
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
