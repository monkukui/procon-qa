<template>
  <div class="questions">
    <QuestionPanel
      title="spritekitを用いたMac向けのアプリでキーボードの同時入力を処理したい"
      body="吾輩は猫である, 名前はまだない"
      questionedTime="2019/05/26 11:24:12"
      :answered="fal"
    />
    <QuestionPanel
      title="for of を使ったReverse String"
      body="世界の果てまでイッテ急"
      questionedTime="2019/05/26 11:24:12"
      :answered="fal"
    />
    <QuestionPanel
      title="AjaxのJavaScriptが動きません"
      body="そこはかとなく先塗るちを"
      questionedTime="2019/05/26 11:24:12"
      :answered="tr"
    />
    <QuestionPanel
      title="TypeScript 型エラー"
      body='string 型で定義した変数に, url query から取得して代入したいです.
            このとき, 型が undefined になる可能性があって (string になるとは限らない) コンパイルエラーとなってしまいます
            どうしたらいいでしょうか??
            '
      questionedTime="2019/05/26 11:24:12"
      :answered="fal"
    />
    <QuestionPanel
      title="TypeScript 型エラー"
      body='string 型で定義した変数に, url query から取得して代入したいです.
            このとき, 型が undefined になる可能性があって (string になるとは限らない) コンパイルエラーとなってしまいます
            どうしたらいいでしょうか??
            '
      questionedTime="2019/05/26 11:24:12"
      :answered="fal"
    />
    <QuestionPanel
      title="TypeScript 型エラー"
      body='string 型で定義した変数に, url query から取得して代入したいです.
            このとき, 型が undefined になる可能性があって (string になるとは限らない) コンパイルエラーとなってしまいます
            どうしたらいいでしょうか??
            '
      questionedTime="2019/05/26 11:24:12"
      :answered="fal"
    />
    <QuestionPanel
      title="TypeScript 型エラー"
      body='string 型で定義した変数に, url query から取得して代入したいです.
            このとき, 型が undefined になる可能性があって (string になるとは限らない) コンパイルエラーとなってしまいます
            どうしたらいいでしょうか??
            '
      questionedTime="2019/05/26 11:24:12"
      :answered="fal"
    />
    <v-pagination
      v-model="curPageId"
      :length="15"
      :total-visible="7"
      circle
    ></v-pagination>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch} from 'vue-property-decorator';
import QuestionPanel from '@/components/QuestionPanel.vue';
import PagiNation from '@/components/PagiNation.vue';


@Component({
  components: {
    QuestionPanel,
    PagiNation,
  },
})
export default class Questions extends Vue {

  // FIXME boolean の変数を作る
  private tr: boolean = true;
  private fal: boolean = false;

  private curPageId: number = 1;
  private changePageId(id: number): void {
    this.curPageId = id;
  }
}
</script>

<style scoped>
.questions {
  margin: 0px 0px 20px-24px;
}
</style>
