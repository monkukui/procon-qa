<template>
  <div class="mypage">
    <v-tabs
      fixed-tabs
      background-color="#607d8d"
      dark
    >
      <v-tab @click="changeMode('profile')">
        プロフィール
      </v-tab>
      <v-tab @click="changeMode('question')">
        質問
      </v-tab>
      <v-tab @click="changeMode('answer')">
        回答
      </v-tab>
      <v-tab @click="changeMode('setting')">
        設定
      </v-tab>
    </v-tabs>

    <div v-if="mode=='profile'">
      <MyProfile />
    </div>
    <div v-if="mode=='question'">
      <MyQuestion />
    </div>
    <div v-if="mode=='answer'">
      <MyAnswer />
    </div>
    <div v-if="mode=='setting'">
      <MySetting />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import UserQuestions from '@/components/UserQuestions.vue';
import UserAnswers from '@/components/UserAnswers.vue';

import MyProfile from '@/components/MyProfile.vue';
import MyQuestion from '@/components/MyQuestion.vue';
import MyAnswer from '@/components/MyAnswer.vue';
import MySetting from '@/components/MySetting.vue';

@Component({
  components: {
    UserQuestions,
    UserAnswers,
    MyProfile,
    MyQuestion,
    MyAnswer,
    MySetting,
  },
})
export default class MyPage extends Vue {
  private mode: string = 'questions';
  private activeBtn: number = 1;
  private userName: string = '';
  private userId: number = 0;
  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.userName = claims.name;
    this.userId = claims.uid;
  };
  private getToken(): any {
    return localStorage.getItem('token');
  };
  private changeMode(mode: string): void {
    this.mode = mode;
  }

}
</script>

<style>
</style>
