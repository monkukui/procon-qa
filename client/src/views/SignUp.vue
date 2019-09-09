<template>
  <div class="login">
    <p>SignUp</p>
    <input v-model="name" />
    <input v-model="password" />
    <button type="button" name="button" @click="signup">signup</button>
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
export default class SignUp extends Vue {
  
  private name: string = '';
  private password: string = '';
  
  private signup(): void {
    const url = '/signup'
    const method = 'POST'
    const headers = {
      'Content-Type': 'application/json; charset=UTF-8'
    }
    const body = JSON.stringify({
      name: this.name,
      password: this.password,
    })

    fetch(url, {method, headers, body}).then(response => {
      if(response.status === 400) {
        alert('Name or Password are empty. Please retry')
      } else if(response.status === 409) {
        alert('Name already exists. Please retry')
        this.name = ''
        this.password = ''
      } else if(response.status === 201) {
        this.$router.push('/login')
      }
    })
  }
}
</script>

<style scoped>
</style>
