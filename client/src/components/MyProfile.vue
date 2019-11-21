<template>
  <div class="myprofile">
    <hr>
    <p>{{ userName }} さん，こんにちは．</p>
    <p>ID は {{ userId }} です</p>
    <p>質問に対するいいねの数は {{ userFavoriteQuestion }} です</p>
    <p>回答に対するいいねの数は {{ userFavoriteAnswer }} です</p>
    <p>獲得したいいねの数の合計は {{ userFavoriteAnswer + userFavoriteQuestion }} です</p>
    <p>合計パート 2 {{ userFavoriteSum }}</p>
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
      console.log(json);
      this.userName = json.name;
      this.userFavoriteQuestion = json.favoriteQuestion;
      this.userFavoriteAnswer = json.favoriteAnswer;
      this.userFavoriteSum = json.favoriteSum;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>
