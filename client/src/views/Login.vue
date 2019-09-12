<template>
  <div class="login">
    <p>LOGIN</p>
    <input v-model="name" />
    <input v-model="password" />
    <button type="button" name="button" @click="login">login</button>
    {{ name }}
    {{ password }}
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
    }
    const body = JSON.stringify({
      name: this.name,
      password: this.password,
    })

    fetch(url, {method, headers, body}).then(response => {
      if(response.ok) {
        return response.json();
      } else {
        alert('Faild to login. Please retry');
        this.name = '';
        this.password = '';
        return {token: ''};
      }
    }).then(json => {
      const token = json.token
      if(token.length > 0) {
        localStorage.setItem('token', token);
        this.$router.push('/');
      }
    })
  }
}
</script>

<style scoped>
</style>
