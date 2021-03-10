<template>
  <div class="user-update"> 
    <v-form
      ref="form"
      v-model="valid"
      lazy-validation
    >
      <v-text-field
        v-model="name"
        :rules="nameRules"
        :counter="14"
        label="ユーザ名"
        required
      >
      </v-text-field>
      <v-text-field
        v-model="twitterId"
        label="Twitter ID"
      >
        <v-icon slot="prepend" color="green">mdi-at</v-icon>
      </v-text-field>

      <v-btn
        color="primary"
        @click="submit"
        :disabled="!valid"
      >
         更新
      </v-btn>
    </v-form>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';

@Component({
  components: {
  },
})

export default class UserUpdateForms extends Vue {

  private name: string = '';
  private twitterId: string = '';
  private valid: boolean = true;
  private nameRules: any = [
    (v: any) => !!v || 'この項目は必須です',
    (v: any) => v.length <= 14 || '14 文字以下にしてください',
  ];

  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    const uid = claims.uid;
    const url = '/api/no-auth/user/' + String(uid);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.name = json.name;
      this.twitterId = json.twitter_id;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
  private submit(): void {
    const url = '/api/update-user';
    const method = 'PUT';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };
    const body = JSON.stringify({
      name: this.name,
      twitter_id: this.twitterId,
    });

    for (let i = 0; i < this.name.length - 1; i++) {
      if ('a' <= this.name[i] && this.name[i] <= 'z') continue;
      if ('A' <= this.name[i] && this.name[i] <= 'Z') continue;
      if ('0' <= this.name[i] && this.name[i] <= '9') continue;
      alert('パスワードは、英字と数字からなる文字列にしてください');
      return;
    }

    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        alert('更新されました');
      } else {
        if (response.status === 400) {
          alert('Name are empty. Please retry');
        } else if (response.status === 409) {
          alert('このユーザ名は使われています．別の名前をお試しください．');
          this.name = '';
        }
      }
    });
  }
}
</script>

<style scoped>
.user-update {
  margin: 50px;
}
</style>
