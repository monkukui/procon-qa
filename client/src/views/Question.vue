<template>
  <div class="question" v-if="exsist==1">
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
  <div v-else-if="exsist==2">
    <h1>質問が見つかりません</h1>
    <p>すでに削除された可能性があります</p>
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
  private qid: any = '';
  private xor: boolean = false;
  private isAnswering: boolean = false;
  private exsist: number = 0; // 0 待機中，1 ok，2 not found

  private clickAnswerButton(): void {
    this.isAnswering = true;
  }

  private changeMode(mode: number): void {
    this.mode = mode;
  }
  private postAnswer(): void {
    this.xor = !this.xor;
  }
  private created(): void {
    // この質問が存在するかを判定

    this.qid = this.$route.query.questionId;
    const url = '/api/no-auth/question/' + this.qid;
    fetch(url).then((response) => {
      if (response.ok) {
        this.exsist = 1;
      } else {
        this.exsist = 2;
      }
    });
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
