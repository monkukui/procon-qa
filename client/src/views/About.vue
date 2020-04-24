<template>
  <div class="mavonEditor">
    <h1>本サイトは作成途中です．</h1>
    <p>ご意見ご要望は
    <v-btn icon href="https://twitter.com/monkukui2">
      <v-icon color="rgb(73, 160, 236)">mdi-twitter</v-icon>
    </v-btn>
    まで．
    <p>
      <v-btn icon href="https://github.com/monkukui/procon-qa">
        <v-icon color="rgb(37, 41, 46)">mdi-github</v-icon>
      </v-btn>
      の issue や PR も募集中.
    </p>

  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({
  components: {
  },
})
export default class About extends Vue {
}
</script>

<style scoped>
</style>
