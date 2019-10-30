<template>
  <div class="login">
    <h1>ログイン</h1>
    <hr>
    <v-container>
      <v-row>
        <v-col
          cols="12"
          md="4"
        >
          <v-text-field
            v-model="name"
            :counter="10"
            label="ユーザ名"
            required
          ></v-text-field>
        </v-col>
        <v-col
          cols="12"
          md="4"
        >
          <v-text-field
            v-model="password"
            :counter="10"
            label="パスワード"
            required
          ></v-text-field>
        </v-col>
      </v-row>
    </v-container>
    <v-btn large color="primary" @click="login">ログイン</v-btn> または <v-btn outlined depressed large to="/signup">新規登録</v-btn>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class Login extends Vue {

  private name: string = '';
  private password: string = '';

  private login(): void {
    const url = '/login';
    const method = 'POST';
    const headers = {
      'Content-Type': 'application/json; charset=UTF-8',
    };
    const body = JSON.stringify({
      name: this.name,
      password: this.password,
    });

    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        return response.json();
      } else {
        alert('Faild to login. Please retry');
        this.name = '';
        this.password = '';
        return {token: ''};
      }
    }).then((json) => {
      const token = json.token;
      if (token.length > 0) {
        localStorage.setItem('token', token);
        this.$router.push('/');
      }
    });
  }
}
</script>

<style scoped>
</style>
