<template>
  <div class="home">
    <div class="main-contents">
      <v-container fluid>
        <v-row>
          <v-col col="12" sm="8">
            <QuestionSetting 
              @click="changeMode"
              :curMode="mode"
            />
            <Questions 
              :mode="mode"
            />
          </v-col>
          <v-col col="12" sm="4">
            <v-card
    max-width="344"
    class="mx-auto"
  >
    <v-list-item>
      <v-list-item-content>
        <v-list-item-title class="headline">Our Changing Planet</v-list-item-title>
        <v-list-item-subtitle>by Kurt Wagner</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>

    <v-img
      src="https://cdn.vuetifyjs.com/images/cards/mountain.jpg"
      height="194"
    ></v-img>

    <v-card-text>
      Visit ten places on our planet that are undergoing the biggest changes today.
    </v-card-text>

    <v-card-actions>
      <v-btn
        text
        color="deep-purple accent-4"
      >
        Read
      </v-btn>
      <v-btn
        text
        color="deep-purple accent-4"
      >
        Bookmark
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-share-variant</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
          </v-col>
        </v-row>
      </v-container>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Questions from '@/components/Questions.vue';
import QuestionSetting from '@/components/QuestionSetting.vue';

@Component({
  components: {
    Questions,
    QuestionSetting,
  },
})
export default class Home extends Vue {
  private mode: number = 1;  // 新着, 回答数, 閲覧数, いいね, 解決済みかどうか
  private changeMode(mode: number): void {
    this.mode = mode;
  }
}
</script>

<style scoped>
.home {
}
</style>
