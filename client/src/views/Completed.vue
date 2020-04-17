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
    <v-row>
      <v-col cols="12" sm="3">
        <v-btn 
          small
          depressed
          to="/"
        >
          ホームに戻る
        </v-btn>
      </v-col>
      <v-col cols="12" sm="3">
        <a
          href="https://twitter.com/share"
          class="twitter-share-button"
          :data-url=shareUrl
          data-text="procon-qa で質問しました！"
          data-size="large"
          data-hashtags="procon-qa"
        >
        </a>
      </v-col>
    </v-row>
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
  private shareUrl: string = '';

  private created(): void {
    // location.reload();
    // location.reload(false);
    this.qid = this.$route.query.qid;
    this.shareUrl = 'https://procon-qa.herokuapp.com/#/question?questionId=' + this.qid;
  }
}
</script>

<style scoped>
.completed {
  width: 50%;
  padding: 20px;
}
.twitter_share{
  max-width: 1000px;
  margin: auto;
}
</style>
