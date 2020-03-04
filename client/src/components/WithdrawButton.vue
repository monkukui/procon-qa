<template>
  <div class="withdraw">
    <v-row justify="center">
      <v-btn
        color="error"
        outlined
        @click.stop="dialog = true"
      >
        退会する
      </v-btn>

      <v-dialog
        v-model="dialog"
        max-width="290"
      >
        <v-card>
          <v-card-title class="headline">
            本当に退会しますか？
          </v-card-title>

          <v-card-text>
            一度退会したら，アカウントを復活させる事はできません．
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn
              color="green darken-1"
              text
              @click="dialog = false"
            >
              NO
            </v-btn>

            <v-btn
              color="green darken-1"
              text
              @click="withdraw"
            >
              YES
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class WithdrawButton extends Vue {

  private userId: number = 0;
  private dialog: boolean = false;
  private withdraw(): void {
    const url = 'api/user/' + String(this.userId);
    const method = 'DELETE';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        // 質問を削除しました
        alert('アカウントを削除しました');
        localStorage.removeItem('token');
        this.$router.push('./');
        location.reload();
      } else {
        alert('不正です');
      }
    });

  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.userId = claims.uid;
  }
}
</script>

<style>
.withdraw {
  margin: 50px;
}
</style>
