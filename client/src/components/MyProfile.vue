<template>
  <div class="myprofile">
    <UserName
      :name="userName"
      :uid="userId"
      size="large"
    />
    <TwitterIcon
      :twitterId="userTwitterId"
      :uid="userId"
      size="100"
      apiSize="o"
    />
    <br>
    <p>Twitter ID： 
      <span v-if="userTwitterId">
        <a :href="userTwitterLink">@{{ userTwitterId }}</a>
      </span>
    </p>
    <p>質問への獲得いいね数： {{ userFavoriteQuestion }}</p>
    <p>回答への獲得いいね数： {{ userFavoriteAnswer }}</p>
    <p>合計獲得いいね数：{{ userFavoriteAnswer + userFavoriteQuestion }}</p>
    <v-divider></v-divider>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import TwitterIcon from '@/components/TwitterIcon.vue';
import UserName from '@/components/UserName.vue';

@Component({
  components: {
    TwitterIcon,
    UserName,
  },
})
export default class MyProfile extends Vue {

  private userName: string = '';
  private userTwitterId: string = '';
  private userTwitterLink: string = '';
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
      this.userTwitterLink = 'https://twitter.com/' + this.userTwitterId;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>
