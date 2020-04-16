<template>
  <div class="completed">
    {{ qid }}
    <v-alert
      outlined
      type="success"
      text
    >
      質問の投稿が完了しました．
    </v-alert>
    <v-btn 
      depressed
      to="/"
    >
      ホームに戻る
    </v-btn>
    <div class="twitter_share">
      <v-btn @click="twitterShare">ツイッターでシェアする</v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';
@Component({
  components: {
  },
})

export default class Completed extends Vue {
  private qid!: any;

  private created(): void {
    this.qid = this.$route.query.qid;
  }

  private twitterShare(): void {
   //シェアする画面を設定
   // TODO 実装
   var shareURL = 'https://twitter.com/intent/tweet?text=' + "procon-qa で質問しました！" + "%20%23proconqa " + '&url=' + 
   "https://procon-qa.herokuapp.com/#/question?questionId=" + this.qid;
   //シェア用の画面へ移行
    location.href = shareURL
  }
}
</script>

<style scoped>
.completed {
  position: absolute;
  top: 50%;
  left: 50%;
  -webkit-transform: translateY(-50%) translateX(-50%);
  transform: translateY(-50%) translateX(-50%);
  padding: 20px;
}
.twitter_share{
  max-width: 1000px;
  margin: auto;
}
</style>
