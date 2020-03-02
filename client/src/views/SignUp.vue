<template>
  <div class="login">
    <h1>新規登録</h1>
    <hr>
    <v-container>
      <v-row>
        <v-col
          cols="12"
          md="6"
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
          md="6"
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
      </v-row>
    </v-container>

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
    });

    if (this.name.length > 14) {
      alert('名前が長すぎます');
      return;
    }

    fetch(url, {method, headers, body}).then((response) => {
      if (response.status === 400) {
        alert('Name or Password are empty. Please retry');
      } else if (response.status === 409) {
        alert('Name already exists. Please retry');
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
