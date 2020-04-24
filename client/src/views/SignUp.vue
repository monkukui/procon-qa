<template>
  <div class="login">
    <v-container>
      <h1>新規登録</h1>
      <hr>
        <v-row>
          <v-col
            cols="12"
            md="8"
          >
            <v-text-field
              v-model="name"
              :counter="14"
              label="ユーザ名"
              required
            ></v-text-field>
          </v-col>
          <v-col
            cols="12"
            md="8"
          >
            <v-text-field
              v-model="twitterId"
              label="TwitterID（任意）"
            ></v-text-field>
          </v-col>
          <v-col
            cols="12"
            md="8"
          >
            <v-text-field
              v-model="password"
              :type="show ? 'text' : 'password'"
              @click:append="show = !show"
              :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
              label="パスワード"
              required
            ></v-text-field>
          </v-col>
          <v-col
            cols="12"
            md="8"
          >
            <v-text-field
              v-model="passwordAgain"
              :type="show ? 'text' : 'password'"
              @click:append="show = !show"
              :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
              label="パスワードの確認"
              required
            ></v-text-field>
          </v-col>
        </v-row>

      <v-alert
        dense
        text
        type="success"
      >
        パスワードは<strong>ハッシュ化</strong>して保存されます．<br>
        ユーザ登録をすれば，質問を投稿したり，質問に回答することができます．<br>
        ユーザ名は，<strong>14 文字以内</strong>にしてください<br>
        新規登録直後に，ログインが要求されます．ご了承ください<br>
      

      </v-alert>
      <v-btn large color="primary" @click="signup">登録</v-btn>
    </v-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class SignUp extends Vue {

  private name: string = '';
  private password: string = '';
  private passwordAgain: string = '';
  private twitterId: string = '';
  private show: boolean = false;

  private signup(): void {
    const url = '/signup';
    const method = 'POST';
    const headers = {
      'Content-Type': 'application/json; charset=UTF-8',
    };
    const body = JSON.stringify({
      name: this.name,
      password: this.password,
      twitter_id: this.twitterId,
    });

    if (this.name.length > 14) {
      alert('ユーザ名が長すぎます');
      return;
    }

    if (this.password !== this.passwordAgain) {
      alert('パスワードが一致しません');
      return;
    }

    fetch(url, {method, headers, body}).then((response) => {
      if (response.status === 400) {
        alert('名前もしくはパスワードが空です');
      } else if (response.status === 409) {
        alert('このユーザ名は使われています．別の名前をお試しください．');
        this.name = '';
        this.password = '';
      } else if (response.status === 201) {
        this.$router.push('/login');
      }
    });
  }
}
</script>

<style scoped>
</style>
