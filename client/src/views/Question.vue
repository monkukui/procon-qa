<template>
  <div class="question">
    <v-container fluid>
      <v-row>
        <v-col cols="12" sm="12">
          <QuestionPanelDetail
            :title="title"
            :body="body"
            :questionedTime="questionedTime"
          />
        </v-col>
          
        <v-divider class="mx-6"></v-divider>
        <v-col cols="12" sm="12" v-if="!isAnswering">
          <span class="answerButton">
            <v-btn x-large color="primary" @click="clickAnswerButton">
              <v-icon left>mdi-pencil</v-icon>回答する
            </v-btn>
          </span>
        </v-col>
        <v-col cols="12" sm="12" v-if="isAnswering">
          <AnswerForm
            @answer="postAnswer"
          />
        </v-col>
        <v-col cols="12" sm="12">
          <AnswerSetting
            @click="changeMode"
            :curMode="mode"
          />
          <Answers 
            :mode="mode" 
            :xor="xor"
          />
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
  private xor: boolean = false;
  private isAnswering: boolean = false;

  private clickAnswerButton(): void {
    this.isAnswering = true;
  }

  private changeMode(mode: number): void {
    this.mode = mode;
  }
  private postAnswer(): void {
    this.xor = !this.xor;
  }
  private mounted() {
    this.title = this.$route.query.title;
    this.body = this.$route.query.body;
    this.questionedTime = this.$route.query.questionedTime;
  }
}
</script>

<style scoped>
.question {
  margin: 40px;
}
.answerButton {
  margin-left: 40px;
}

</style>
