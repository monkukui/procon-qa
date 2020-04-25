<template>
  <div class="myprofile">
    <h1>{{ userName }}</h1>
    <TwitterIcon
      :twitterId="userTwitterId"
      :uid="userId"
      size="100"
      apiSize="o"
    />
    
    <p>質問への獲得いいね数： {{ userFavoriteQuestion }}</p>
    <p>回答への獲得いいね数： {{ userFavoriteAnswer }}</p>
    <p>合計獲得いいね数：{{ userFavoriteAnswer + userFavoriteQuestion }}</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import TwitterIcon from '@/components/TwitterIcon.vue';

@Component({
  components: {
    TwitterIcon,
  },
})
export default class MyProfile extends Vue {

  private userName: string = '';
  private userTwitterId: string = '';
  private userId: string | Array<(string | null)> = '';
  private userFavoriteAnswer: number = 0;
  private userFavoriteQuestion: number = 0;
  private userFavoriteSum: number = 0;

  private created(): void {
    this.userId = this.$route.query.uid;
    this.setUser();
  }
  private setUser(): void {
    const url = 'api/no-auth/user/' + String(this.userId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userName = json.name;
      this.userTwitterId = json.twitter_id;
      this.userFavoriteQuestion = json.favorite_question;
      this.userFavoriteAnswer = json.favorite_answer;
      this.userFavoriteSum = json.favorite_sum;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>
