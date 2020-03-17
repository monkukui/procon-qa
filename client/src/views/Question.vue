<template>
  <div class="question">

    <v-container fluid>
      <v-row>
        <v-col cols="12" sm="8">
          <QuestionPanelDetail
            :title="title"
            :body="body"
            :questionedTime="questionedTime"
          />
        </v-col>
        <v-col cols="12" sm="8">
          <AnswerSetting
            @click="changeMode"
            :curMode="mode"
          />
          <Answers 
            :mode="mode" 
          />
        </v-col>
        <v-col cols="12" sm="8">
          <AnswerForm />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QuestionPanelDetail from '@/components/QuestionPanelDetail.vue';
import Answers from '@/components/Answers.vue';
import AnswerForm from '@/components/AnswerForm.vue';
import AnswerSetting from '@/components/AnswerSetting.vue';


@Component({
  components: {
    QuestionPanelDetail,
    Answers,
    AnswerForm,
    AnswerSetting,
  },
})
export default class Question extends Vue {
  private mode: number = 1;  // 新着(1)，いいね(2)
  // TODO getParam とかでとってくる
  // TODO any は最悪なのでなんとかする (string | undefined とかにしてもエラーがでる)
  private title: any = '';
  private body: any = '';
  private questionedTime: any = '';

  private changeMode(mode: number): void {
    this.mode = mode;
  }
  private mounted() {
    this.title = this.$route.query.title;
    this.body = this.$route.query.body;
    this.questionedTime = this.$route.query.questionedTime;
  }
}
</script>
