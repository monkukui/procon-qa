<template>
  <div class="mypage">
    <h1>{{ userName }} id : {{ userId }}</h1>
    <UserQuestions />
    <UserAnswers />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import UserQuestions from '@/components/UserQuestions.vue';
import UserAnswers from '@/components/UserAnswers.vue';

@Component({
  components: {
    UserQuestions,
    UserAnswers,
  },
})
export default class MyPage extends Vue {
  private userName: string = '';
  private userId: number = 3;
  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.userName = claims.name;
    this.userId = claims.uid;
  };
  private getToken(): any {
    return localStorage.getItem('token');
  };

}
</script>

<style>
</style>
