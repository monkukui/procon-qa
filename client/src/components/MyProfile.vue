<template>
  <div class="myprofile">
    <hr>
    <p>{{ userName }} さんの獲得いいね情報．</p>
    <v-divider />
    <p>質問への獲得いいね数： {{ userFavoriteQuestion }}</p>
    <p>回答への獲得いいね数： {{ userFavoriteAnswer }}</p>
    <p>合計獲得いいね数：{{ userFavoriteAnswer + userFavoriteQuestion }}</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';

@Component({
})
export default class MyProfile extends Vue {

  private userName: string = '';
  private userId: number = 0;
  private userFavoriteAnswer: number = 0;
  private userFavoriteQuestion: number = 0;
  private userFavoriteSum: number = 0;

  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.userName = claims.name;
    this.userId = claims.uid;
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
