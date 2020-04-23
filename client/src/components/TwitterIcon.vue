<template>
  <div class="twitter_icon">

    <!--v-avatar color="indigo">
      <v-icon dark>mdi-account-circle</v-icon>
    </v-avatar-->
    <v-avatar :size=size>
      <img
        :src=img_url
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
  private api_size!: string;  // m (24 * 24), n (48 * 48), b (73 * 73), o (original size)
  @Prop()
  private twitter_id!: string; 

  private img_url: string = '';

  private created() {
    this.img_url = "https://tweetimag.es/i/" + this.twitter_id + "_" + this.api_size;
  }
}
</script>

<style scoped>
.twitter_icon {
  margin: 5px;
}
</style>
