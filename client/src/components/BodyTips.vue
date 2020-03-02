<template>
  <div>
    <v-alert
      border="top"
      colored-border
      type="info"
      elevation="2"
    >
      <h3>本文に関する tips</h3>
      <v-divider/>
      <br>
      <p>バッククォート`でコードフェンスを作ることができます</p>
      <p>(例)</p>
      <p>```C</p>
      <p>int main() {</p>
      <p>&nbsp;&nbsp;printf("hello world\n");</p>
      <p>&nbsp;&nbsp;return 0;</p>
      <p>}</p>
      <p>```</p>
      <p><em>斜体</em>や<strong>太字</strong>にできます</p>
      <p>その他マークダウンの記法がそのまま使えます</p>
    </v-alert>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class BodyTips extends Vue {
}
</script>

<style>
</style>
