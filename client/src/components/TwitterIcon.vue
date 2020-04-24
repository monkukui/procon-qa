<template>
  <div class="twitter_icon">

    <!--v-avatar color="indigo">
      <v-icon dark>mdi-account-circle</v-icon>
    </v-avatar-->
    <v-avatar :size=size>
      <img
        :src=imgUrl
        alt="icon"
      >
    </v-avatar>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class TwitterIcon extends Vue {
  @Prop()
  private size!: string;
  @Prop()
  private apiSize!: string;  // m (24 * 24), n (48 * 48), b (73 * 73), o (original size)
  @Prop()
  private twitterId!: string;

  private imgUrl: string = '';

  private created() {
    this.imgUrl = 'https://tweetimag.es/i/' + this.twitterId + '_' + this.apiSize;
  }
}
</script>

<style scoped>
.twitter_icon {
  margin: 5px;
}
</style>
