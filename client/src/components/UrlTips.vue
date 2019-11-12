<template>
  <div>
    <v-alert
      border="top"
      colored-border
      type="info"
      elevation="2"
    >
      <h3>URL に関する tips</h3>
      <v-divider/>
      <br>
      <p>自分の提出のリンクを記載すると，回答がつきやすくなります.</p>
      <p>(AtCoder): /submissions/(提出 ID)</p>
      <p>(注意): リンクが submissions/me になってないことを確認してください</p> 
    </v-alert>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class UrlTips extends Vue {
}
</script>

<style>
</style>
