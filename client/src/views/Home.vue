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
            <SiteInfomation />
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
import SiteInfomation from '@/components/SiteInfomation.vue';

@Component({
  components: {
    Questions,
    QuestionSetting,
    SiteInfomation,
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
