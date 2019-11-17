<template>
  <div>
    <v-alert
      border="top"
      colored-border
      type="info"
      elevation="2"
      transition="scale-transition"
    >
      <h3>タイトルに関する tips</h3>
      <v-divider/>
      <br>
      <p>このサイトは<u>競技プログラミング</u>に関する Q & A サイトです</p>
      <p>質問内容をより具体的にすれば，良い回答がつきやすくなります</p>
      <p><span style="color: #e53935; font-size: large">(悪い例):</span>自分のコードが WA になる理由がわかりません</p> 
      <p><span style="color: #0d47a1; font-size: large;">(良い例):</span>ABC040 C 問題について，自分の貪欲法が WA になる理由がわかりません</p>
    </v-alert>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class TitleTips extends Vue {
}
</script>

<style>
</style>
